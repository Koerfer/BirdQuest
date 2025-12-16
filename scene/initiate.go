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

func initiateObjects(path string) {
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

	prepareCollisionObjects(&jMap, "BackgroundCollisions", path, jsonCollisionsGidStart)
	prepareItems(&jMap, "Items", path, jsonItemsGidStart)
	prepareBloons(&jMap, "Items", path, jsonBloonsGidStart)
	prepareCollisions(&jMap)
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

func prepareItems(jMap *jsonMap, layerName, path string, startId int) {
	itemSpritesRaw := rl.LoadTexture(filepath.Join(path, "item_sprites.png"))
	jsonItemContents, err := os.ReadFile(filepath.Join(path, "items.tsj"))
	if err != nil {
		log.Fatal(err)
	}
	var jsonItems jsonSprites
	err = json.Unmarshal(jsonItemContents, &jsonItems)
	if err != nil {
		log.Fatal(err)
	}

	itemSprites := prepareSprites(itemSpritesRaw, &jsonItems, startId)
	CurrentScene.ItemObjects = &Items{
		Texture: itemSpritesRaw,
		Objects: make([]*Object, 0),
	}

	for _, layer := range jMap.Layers {
		if layer.Name != layerName {
			continue
		}

		for i, val := range layer.Data {
			prepareItem(i, val, startId, jsonItems.TileCount, itemSprites)
		}
	}
}

func prepareBloons(jMap *jsonMap, layerName, path string, startId int) {
	bloonSpritesRaw := rl.LoadTexture(filepath.Join(path, "bloons.png"))
	jsonBloonContents, err := os.ReadFile(filepath.Join(path, "bloons.tsj"))
	if err != nil {
		log.Fatal(err)
	}
	var jsonBloons jsonSprites
	err = json.Unmarshal(jsonBloonContents, &jsonBloons)
	if err != nil {
		log.Fatal(err)
	}

	bloonSprites := prepareSprites(bloonSpritesRaw, &jsonBloons, startId)
	CurrentScene.Bloons = &Bloons{
		Texture:      bloonSpritesRaw,
		BloonObjects: make([]*Bloon, 0),
	}

	for _, layer := range jMap.Layers {
		if layer.Name != layerName {
			continue
		}

		for i, val := range layer.Data {
			prepareBloon(i, val, startId, jsonBloons.TileCount, bloonSprites)
		}
	}
}

func prepareBloon(i, val, startId, n int, sprites *sprites) {
	if val == 0 {
		return
	}
	if val < startId {
		return
	}
	if val >= startId+n {
		return
	}

	x := float32(i % CurrentScene.WidthInTiles * global.TileWidth)
	y := float32(i / CurrentScene.WidthInTiles * global.TileWidth)

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

			CurrentScene.Bloons.BloonObjects = append(CurrentScene.Bloons.BloonObjects, bloon)
			return
		}
	}
}

func prepareItem(i, val, startId, n int, sprites *sprites) {
	if val == 0 {
		return
	}
	if val < startId {
		return
	}
	if val >= startId+n {
		return
	}

	x := float32(i % CurrentScene.WidthInTiles * global.TileWidth)
	y := float32(i / CurrentScene.WidthInTiles * global.TileWidth)

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

			CurrentScene.ItemObjects.Objects = append(CurrentScene.ItemObjects.Objects, object)
			return
		}
	}
}

func prepareCollisionObjects(jMap *jsonMap, layerName, path string, startId int) {
	collisionSpritesRaw := rl.LoadTexture(filepath.Join(path, "collision_sprites.png"))
	jsonCollisionContents, err := os.ReadFile(filepath.Join(path, "collisions.tsj"))
	if err != nil {
		log.Fatal(err)
	}

	var jsonCollisions jsonSprites
	err = json.Unmarshal(jsonCollisionContents, &jsonCollisions)
	if err != nil {
		log.Fatal(err)
	}

	collisionSprites := prepareSprites(collisionSpritesRaw, &jsonCollisions, startId)
	CurrentScene.CollisionObjects = &CollisionItems{
		Texture:     collisionSpritesRaw,
		DrawFirst:   make([]*Object, 0),
		DrawDynamic: make([]*Object, 0),
		DrawLast:    make([]*Object, 0),
	}

	for _, layer := range jMap.Layers {
		if layer.Name != layerName {
			continue
		}

		for i, val := range layer.Data {
			prepareCollisionObject(i, val, startId, jsonCollisions.TileCount, collisionSprites)
		}
	}
}

func prepareCollisionObject(i, val, startId, n int, sprites *sprites) {
	if val == 0 {
		return
	}
	if val < startId {
		return
	}
	if val >= startId+n {
		return
	}

	x := float32(i % CurrentScene.WidthInTiles * global.TileWidth)
	y := float32(i / CurrentScene.WidthInTiles * global.TileWidth)

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
				CurrentScene.CollisionObjects.DrawFirst = append(CurrentScene.CollisionObjects.DrawFirst, object)
			} else if prop.AlwaysRenderLast {
				CurrentScene.CollisionObjects.DrawLast = append(CurrentScene.CollisionObjects.DrawLast, object)
			} else {
				CurrentScene.CollisionObjects.DrawDynamic = append(CurrentScene.CollisionObjects.DrawDynamic, object)
			}
			return
		}
	}
}
