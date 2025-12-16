package scene

import (
	"BirdQuest/global"
	"encoding/json"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
	"os"
	"path/filepath"
)

type JsonMap struct {
	Layers   []*JsonLayer   `json:"layers"`
	TileSets []*JsonTileSet `json:"tilesets"`
}

type JsonTileSet struct {
	FirstGid int    `json:"firstgid"`
	Source   string `json:"source"`
}

type JsonLayer struct {
	Data    []int         `json:"data"`
	Objects []*JsonObject `json:"objects"`
	Width   int           `json:"width"`
	Height  int           `json:"height"`
	Name    string        `json:"name"`
}

type JsonObject struct {
	Height float32 `json:"height"`
	Width  float32 `json:"width"`
	X      float32 `json:"x"`
	Y      float32 `json:"y"`
}

type JsonSprites struct {
	TileCount   int          `json:"tilecount"`
	ImageHeight int          `json:"imageheight"`
	ImageWidth  int          `json:"imagewidth"`
	TileWidth   int          `json:"tilewidth"`
	TileHeight  int          `json:"tileheight"`
	Tiles       []*JsonTiles `json:"tiles"`
}

type JsonProperty struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value int    `json:"value"`
}

type JsonTiles struct {
	Id         int             `json:"id"`
	Properties []*JsonProperty `json:"properties"`
}

func InitiateObjects(path string) ([]*Object, []*rl.Rectangle, []*Object, []*Bloon, *Player) {

	itemSpritesRaw := rl.LoadTexture(filepath.Join(path, "item_sprites.png"))
	collisionSpritesRaw := rl.LoadTexture(filepath.Join(path, "collision_sprites.png"))
	chiliAnimationsRaw := rl.LoadTexture(filepath.Join(path, "chili.png"))
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

	var jsonMap JsonMap
	err = json.Unmarshal(jsonMapContents, &jsonMap)
	if err != nil {
		log.Fatal(err)
	}

	var jsonItems JsonSprites
	var jsonItemsGidStart int
	err = json.Unmarshal(jsonItemsContents, &jsonItems)
	if err != nil {
		log.Fatal(err)
	}

	var jsonBloons JsonSprites
	var jsonBloonsGidStart int
	err = json.Unmarshal(jsonBloonsContents, &jsonBloons)
	if err != nil {
		log.Fatal(err)
	}

	var jsonCollisions JsonSprites
	var jsonCollisionsGidStart int
	err = json.Unmarshal(jsonCollisionContents, &jsonCollisions)
	if err != nil {
		log.Fatal(err)
	}

	for _, tileSet := range jsonMap.TileSets {
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

	chiliAnimations := &Sprites{
		Texture:      chiliAnimationsRaw,
		TileWidth:    jsonItems.TileWidth,
		TileHeight:   jsonItems.TileHeight,
		WidthInTiles: int(chiliAnimationsRaw.Width) / jsonItems.TileWidth,
	}

	itemObjects := prepareObjects(&jsonMap, itemsSprites, "Items", jsonItemsGidStart, jsonItems.TileCount)
	collisionObjects3d := prepareObjects(&jsonMap, collisionSprites, "BackgroundCollisions", jsonCollisionsGidStart, jsonCollisions.TileCount)
	collisionObjects := prepareCollisions(&jsonMap)
	bloonsObjects := prepareBloons(&jsonMap, bloonsSprites, "Items", jsonBloonsGidStart, jsonBloons.TileCount)
	player := preparePlayer(chiliAnimations)

	return itemObjects, collisionObjects, collisionObjects3d, bloonsObjects, player
}

func prepareSprites(rawSprites rl.Texture2D, jsonObject *JsonSprites, startGid int) *Sprites {
	sprites := &Sprites{
		Texture:      rawSprites,
		TileWidth:    jsonObject.TileWidth,
		TileHeight:   jsonObject.TileHeight,
		WidthInTiles: int(rawSprites.Width) / jsonObject.TileWidth,
	}
	sprites.Properties = []*Property{}
	for _, tile := range jsonObject.Tiles {
		property := &Property{id: tile.Id + startGid}
		for _, jsonProperty := range tile.Properties {
			switch jsonProperty.Name {
			case "HitBoxX":
				property.HitBoxX = jsonProperty.Value
			case "HitBoxY":
				property.HitBoxY = jsonProperty.Value
			case "HitBoxWidth":
				property.HitBoxWidth = jsonProperty.Value
			case "HitBoxHeight":
				property.HitBoxHeight = jsonProperty.Value
			case "AlwaysAfter":
				switch jsonProperty.Value {
				case 1:
					property.AlwaysRenderLast = true
				default:
					property.AlwaysRenderLast = false
				}
			case "AlwaysFirst":
				switch jsonProperty.Value {
				case 1:
					property.AlwaysRenderFirst = true
				default:
					property.AlwaysRenderFirst = false
				}
			}
		}
		sprites.Properties = append(sprites.Properties, property)
	}

	return sprites
}

func prepareObjects(jsonMap *JsonMap, sprites *Sprites, layerName string, startId, n int) []*Object {
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

func prepareBloons(jsonMap *JsonMap, sprites *Sprites, layerName string, startId, n int) []*Bloon {
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

func prepareObject(i, val, startId, n int, sprites *Sprites) *Object {
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
		Rectangle: sprites.GetSrc(val - startId),
	}

	for _, property := range sprites.Properties {
		if property.id == val {
			object.HitBox = rl.Rectangle{
				X:      (x + float32(property.HitBoxX)) * global.VariableSet.EntityScale,
				Y:      (y + float32(property.HitBoxY)) * global.VariableSet.EntityScale,
				Width:  float32(property.HitBoxWidth) * global.VariableSet.EntityScale,
				Height: float32(property.HitBoxHeight) * global.VariableSet.EntityScale,
			}
			object.AlwaysRenderLast = property.AlwaysRenderLast
			object.AlwaysRenderFirst = property.AlwaysRenderFirst
			break
		}
	}

	return object
}
