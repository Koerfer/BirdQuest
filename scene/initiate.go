package scene

import (
	"BirdQuest/global"
	"encoding/json"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
	"os"
	"path/filepath"
)

type jsonMap struct {
	Layers   []*jsonLayer   `json:"layers"`
	TileSets []*jsonTileSet `json:"tilesets"`
}

type jsonTileSet struct {
	FirstGid int    `json:"firstgid"`
	Source   string `json:"source"`
}

type jsonLayer struct {
	Data    []int         `json:"data"`
	Objects []*jsonObject `json:"objects"`
	Width   int           `json:"width"`
	Height  int           `json:"height"`
	Name    string        `json:"name"`
}

type jsonObject struct {
	Height float32 `json:"height"`
	Width  float32 `json:"width"`
	X      float32 `json:"x"`
	Y      float32 `json:"y"`
}

type jsonSprites struct {
	TileCount   int          `json:"tilecount"`
	ImageHeight int          `json:"imageheight"`
	ImageWidth  int          `json:"imagewidth"`
	TileWidth   int          `json:"tilewidth"`
	TileHeight  int          `json:"tileheight"`
	Tiles       []*jsonTiles `json:"tiles"`
}

type jsonProperty struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value int    `json:"value"`
}

type jsonTiles struct {
	Id         int             `json:"id"`
	Properties []*jsonProperty `json:"properties"`
}

func initiateObjects(path string) ([]*Object, []*rl.Rectangle, []*Object, []*Bloon) {
	itemSpritesRaw := rl.LoadTexture(filepath.Join(path, "item_sprites.png"))
	collisionSpritesRaw := rl.LoadTexture(filepath.Join(path, "collision_sprites.png"))
	bloonsSpritesRaw := rl.LoadTexture(filepath.Join(path, "bloons.png"))

	jsonMapContents, err := os.ReadFile(filepath.Join(path, "map.tmj"))
	if err != nil {
		log.Fatal(err)
	}
	jsonItemsContents, err := os.ReadFile(filepath.Join(path, "items.tsj"))
	if err != nil {
		log.Fatal(err)
	}
	jsonBloonsContents, err := os.ReadFile(filepath.Join(path, "bloons.tsj"))
	if err != nil {
		log.Fatal(err)
	}
	jsonCollisionContents, err := os.ReadFile(filepath.Join(path, "collisions.tsj"))
	if err != nil {
		log.Fatal(err)
	}

	var jMap jsonMap
	err = json.Unmarshal(jsonMapContents, &jMap)
	if err != nil {
		log.Fatal(err)
	}

	var jsonItems jsonSprites
	var jsonItemsGidStart int
	err = json.Unmarshal(jsonItemsContents, &jsonItems)
	if err != nil {
		log.Fatal(err)
	}

	var jsonBloons jsonSprites
	var jsonBloonsGidStart int
	err = json.Unmarshal(jsonBloonsContents, &jsonBloons)
	if err != nil {
		log.Fatal(err)
	}

	var jsonCollisions jsonSprites
	var jsonCollisionsGidStart int
	err = json.Unmarshal(jsonCollisionContents, &jsonCollisions)
	if err != nil {
		log.Fatal(err)
	}

	for _, tileSet := range jMap.TileSets {
		switch tileSet.Source {
		case "items.tsj":
			jsonItemsGidStart = tileSet.FirstGid
		case "bloons.tsj":
			jsonBloonsGidStart = tileSet.FirstGid
		case "collisions.tsj":
			jsonCollisionsGidStart = tileSet.FirstGid
		}
	}

	itemsSprites := prepareSprites(itemSpritesRaw, &jsonItems, jsonItemsGidStart)
	bloonsSprites := prepareSprites(bloonsSpritesRaw, &jsonBloons, jsonBloonsGidStart)
	collisionSprites := prepareSprites(collisionSpritesRaw, &jsonCollisions, jsonCollisionsGidStart)

	itemObjects := prepareObjects(&jMap, itemsSprites, "Items", jsonItemsGidStart, jsonItems.TileCount)
	collisionObjects := prepareObjects(&jMap, collisionSprites, "BackgroundCollisions", jsonCollisionsGidStart, jsonCollisions.TileCount)
	collisionBoxes := prepareCollisions(&jMap)
	bloonsObjects := prepareBloons(&jMap, bloonsSprites, "Items", jsonBloonsGidStart, jsonBloons.TileCount)

	return itemObjects, collisionBoxes, collisionObjects, bloonsObjects
}

func prepareSprites(rawSprites rl.Texture2D, jsonObject *jsonSprites, startGid int) *sprites {
	ss := &sprites{
		Texture:      rawSprites,
		TileWidth:    jsonObject.TileWidth,
		TileHeight:   jsonObject.TileHeight,
		WidthInTiles: int(rawSprites.Width) / jsonObject.TileWidth,
	}
	ss.Properties = []*property{}
	for _, tile := range jsonObject.Tiles {
		prop := &property{id: tile.Id + startGid}
		for _, jProperty := range tile.Properties {
			switch jProperty.Name {
			case "HitBoxX":
				prop.HitBoxX = jProperty.Value
			case "HitBoxY":
				prop.HitBoxY = jProperty.Value
			case "HitBoxWidth":
				prop.HitBoxWidth = jProperty.Value
			case "HitBoxHeight":
				prop.HitBoxHeight = jProperty.Value
			case "AlwaysAfter":
				switch jProperty.Value {
				case 1:
					prop.AlwaysRenderLast = true
				default:
					prop.AlwaysRenderLast = false
				}
			case "AlwaysFirst":
				switch jProperty.Value {
				case 1:
					prop.AlwaysRenderFirst = true
				default:
					prop.AlwaysRenderFirst = false
				}
			}
		}
		ss.Properties = append(ss.Properties, prop)
	}

	return ss
}

func prepareObjects(jsonMap *jsonMap, sprites *sprites, layerName string, startId, n int) []*Object {
	objects := make([]*Object, 0)

	for _, layer := range jsonMap.Layers {
		if layer.Name != layerName {
			continue
		}

		for i, val := range layer.Data {
			object := prepareObject(i, val, startId, n, sprites)
			if object == nil {
				continue
			}

			objects = append(objects, object)
		}
	}

	return objects
}

func prepareBloons(jsonMap *jsonMap, sprites *sprites, layerName string, startId, n int) []*Bloon {
	bloons := make([]*Bloon, 0)

	for _, layer := range jsonMap.Layers {
		if layer.Name != layerName {
			continue
		}

		for i, val := range layer.Data {
			object := prepareObject(i, val, startId, n, sprites)
			if object == nil {
				continue
			}

			bloon := &Bloon{
				Lives: startId + 3 - val,
			}
			bloon.Object = *object

			bloons = append(bloons, bloon)
		}
	}

	return bloons
}

func prepareObject(i, val, startId, n int, sprites *sprites) *Object {
	if val == 0 {
		return nil
	}
	if val < startId {
		return nil
	}
	if val >= startId+n {
		return nil
	}

	x := float32(i % CurrentScene.WidthInTiles * global.TileWidth)
	y := float32(i / CurrentScene.WidthInTiles * global.TileWidth)

	object := &Object{
		Position:  rl.Vector2{X: x * global.VariableSet.EntityScale, Y: y * global.VariableSet.EntityScale},
		Texture:   sprites.Texture,
		Rectangle: sprites.GetRectangleAreaInTexture(val - startId),
	}

	for _, prop := range sprites.Properties {
		if prop.id == val {
			object.HitBox = rl.Rectangle{
				X:      (x + float32(prop.HitBoxX)) * global.VariableSet.EntityScale,
				Y:      (y + float32(prop.HitBoxY)) * global.VariableSet.EntityScale,
				Width:  float32(prop.HitBoxWidth) * global.VariableSet.EntityScale,
				Height: float32(prop.HitBoxHeight) * global.VariableSet.EntityScale,
			}
			object.AlwaysRenderLast = prop.AlwaysRenderLast
			object.AlwaysRenderFirst = prop.AlwaysRenderFirst
			break
		}
	}

	return object
}
