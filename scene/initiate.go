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
	Height     float32         `json:"height"`
	Width      float32         `json:"width"`
	X          float32         `json:"x"`
	Y          float32         `json:"y"`
	Properties []*jsonProperty `json:"properties"`
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
	Value any    `json:"value"`
}

type jsonTiles struct {
	Id         int             `json:"id"`
	Properties []*jsonProperty `json:"properties"`
}

func initiateObjects(path string, scene *Scene) {
	jsonMapContents, err := os.ReadFile(filepath.Join(path, "map.tmj"))
	if err != nil {
		log.Fatal(err)
	}

	var jMap jsonMap
	err = json.Unmarshal(jsonMapContents, &jMap)
	if err != nil {
		log.Fatal(err)
	}

	var jsonCollisionsGidStart int
	var jsonItemsGidStart int
	var jsonBloonsGidStart int
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

	prepareCollisionObjects(&jMap, "BackgroundCollisions", path, jsonCollisionsGidStart, scene)
	prepareItems(&jMap, "Items", path, jsonItemsGidStart, scene)
	prepareBloons(&jMap, "Items", path, jsonBloonsGidStart, scene)
	prepareCollisions(&jMap, scene)
	prepareDoors(&jMap, scene)
}

func prepareSprites(rawSprites rl.Texture2D, jsonObject *jsonSprites, startGid int) *sprites {
	ss := &sprites{
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
				prop.HitBoxX = int(jProperty.Value.(float64))
			case "HitBoxY":
				prop.HitBoxY = int(jProperty.Value.(float64))
			case "HitBoxWidth":
				prop.HitBoxWidth = int(jProperty.Value.(float64))
			case "HitBoxHeight":
				prop.HitBoxHeight = int(jProperty.Value.(float64))
			case "AlwaysAfter":
				switch jProperty.Value.(float64) {
				case 1:
					prop.AlwaysRenderLast = true
				default:
					prop.AlwaysRenderLast = false
				}
			case "AlwaysFirst":
				switch jProperty.Value.(float64) {
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

func prepareItems(jMap *jsonMap, layerName, path string, startId int, scene *Scene) {
	jsonItemContents, err := os.ReadFile(filepath.Join(path, "items.tsj"))
	if err != nil {
		log.Fatal(err)
	}
	var jsonItems jsonSprites
	err = json.Unmarshal(jsonItemContents, &jsonItems)
	if err != nil {
		log.Fatal(err)
	}

	scene.ItemObjects = &Items{
		Objects: make([]*Object, 0),
	}

	itemSpritesRaw := prepareTexture(path, "item_sprites.png")
	scene.ItemObjects.Texture = itemSpritesRaw
	itemSprites := prepareSprites(itemSpritesRaw, &jsonItems, startId)

	for _, layer := range jMap.Layers {
		if layer.Name != layerName {
			continue
		}

		for i, val := range layer.Data {
			prepareItem(i, val, startId, jsonItems.TileCount, itemSprites, scene)
		}
	}
}

func prepareBloons(jMap *jsonMap, layerName, path string, startId int, scene *Scene) {
	jsonBloonContents, err := os.ReadFile(filepath.Join(path, "bloons.tsj"))
	if err != nil {
		log.Fatal(err)
	}
	var jsonBloons jsonSprites
	err = json.Unmarshal(jsonBloonContents, &jsonBloons)
	if err != nil {
		log.Fatal(err)
	}

	scene.Bloons = &Bloons{
		BloonObjects: make([]*Bloon, 0),
	}

	bloonSpritesRaw := prepareTexture(path, "bloons.png")
	scene.Bloons.Texture = bloonSpritesRaw
	bloonSprites := prepareSprites(bloonSpritesRaw, &jsonBloons, startId)

	for _, layer := range jMap.Layers {
		if layer.Name != layerName {
			continue
		}

		for i, val := range layer.Data {
			prepareBloon(i, val, startId, jsonBloons.TileCount, bloonSprites, scene)
		}
	}
}

func prepareBloon(i, val, startId, n int, sprites *sprites, scene *Scene) {
	if val == 0 {
		return
	}
	if val < startId {
		return
	}
	if val >= startId+n {
		return
	}

	x := float32(i % scene.WidthInTiles * global.TileWidth)
	y := float32(i / scene.WidthInTiles * global.TileWidth)

	bloonObject := &Object{
		Position:  rl.Vector2{X: x * global.VariableSet.EntityScale, Y: y * global.VariableSet.EntityScale},
		Rectangle: sprites.GetRectangleAreaInTexture(val - startId),
	}

	for _, prop := range sprites.Properties {
		if prop.id == val {
			bloonObject.HitBox = rl.Rectangle{
				X:      (x + float32(prop.HitBoxX)) * global.VariableSet.EntityScale,
				Y:      (y + float32(prop.HitBoxY)) * global.VariableSet.EntityScale,
				Width:  float32(prop.HitBoxWidth) * global.VariableSet.EntityScale,
				Height: float32(prop.HitBoxHeight) * global.VariableSet.EntityScale,
			}

			bloon := &Bloon{
				Lives: startId + 3 - val,
			}
			bloon.Object = *bloonObject

			scene.Bloons.BloonObjects = append(scene.Bloons.BloonObjects, bloon)
			return
		}
	}
}

func prepareItem(i, val, startId, n int, sprites *sprites, scene *Scene) {
	if val == 0 {
		return
	}
	if val < startId {
		return
	}
	if val >= startId+n {
		return
	}

	x := float32(i % scene.WidthInTiles * global.TileWidth)
	y := float32(i / scene.WidthInTiles * global.TileWidth)

	object := &Object{
		Position:  rl.Vector2{X: x * global.VariableSet.EntityScale, Y: y * global.VariableSet.EntityScale},
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

			scene.ItemObjects.Objects = append(scene.ItemObjects.Objects, object)
			return
		}
	}
}

func prepareCollisionObjects(jMap *jsonMap, layerName, path string, startId int, scene *Scene) {
	jsonCollisionContents, err := os.ReadFile(filepath.Join(path, "collisions.tsj"))
	if err != nil {
		log.Fatal(err)
	}

	var jsonCollisions jsonSprites
	err = json.Unmarshal(jsonCollisionContents, &jsonCollisions)
	if err != nil {
		log.Fatal(err)
	}

	scene.CollisionObjects = &CollisionItems{
		DrawFirst:   make([]*Object, 0),
		DrawDynamic: make([]*Object, 0),
		DrawLast:    make([]*Object, 0),
	}

	collisionSpritesRaw := prepareTexture(path, "collision_sprites.png")
	scene.CollisionObjects.Texture = collisionSpritesRaw
	collisionSprites := prepareSprites(collisionSpritesRaw, &jsonCollisions, startId)

	for _, layer := range jMap.Layers {
		if layer.Name != layerName {
			continue
		}

		for i, val := range layer.Data {
			prepareCollisionObject(i, val, startId, jsonCollisions.TileCount, collisionSprites, scene)
		}
	}
}

func prepareCollisionObject(i, val, startId, n int, sprites *sprites, scene *Scene) {
	if val == 0 {
		return
	}
	if val < startId {
		return
	}
	if val >= startId+n {
		return
	}

	x := float32(i % scene.WidthInTiles * global.TileWidth)
	y := float32(i / scene.WidthInTiles * global.TileWidth)

	object := &Object{
		Position:  rl.Vector2{X: x * global.VariableSet.EntityScale, Y: y * global.VariableSet.EntityScale},
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

			if prop.AlwaysRenderFirst {
				scene.CollisionObjects.DrawFirst = append(scene.CollisionObjects.DrawFirst, object)
			} else if prop.AlwaysRenderLast {
				scene.CollisionObjects.DrawLast = append(scene.CollisionObjects.DrawLast, object)
			} else {
				scene.CollisionObjects.DrawDynamic = append(scene.CollisionObjects.DrawDynamic, object)
			}
			return
		}
	}
}

func prepareTexture(path, pngName string) rl.Texture2D {
	return rl.LoadTexture(filepath.Join(path, pngName))
}
