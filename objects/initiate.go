package objects

import (
	"BirdQuest/global"
	"encoding/json"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
	"os"
	"path/filepath"
	"time"
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
	Data   []int `json:"data"`
	Width  int   `json:"width"`
	Height int   `json:"height"`
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

func InitiateObjects(cwd string, collisionSpritesRaw, itemSpritesRaw, bloonsSpritesRaw, chiliAnimationsRaw rl.Texture2D) ([]*Object, []*Object, []*Bloon, *Player) {

	jsonMapContents, err := os.ReadFile(filepath.Join(cwd, "sprites/map.tmj"))
	if err != nil {
		log.Fatal(err)
	}
	jsonCollisionsContents, err := os.ReadFile(filepath.Join(cwd, "sprites/collisions.tsj"))
	if err != nil {
		log.Fatal(err)
	}
	jsonItemsContents, err := os.ReadFile(filepath.Join(cwd, "sprites/items.tsj"))
	if err != nil {
		log.Fatal(err)
	}
	jsonBloonsContents, err := os.ReadFile(filepath.Join(cwd, "sprites/bloons.tsj"))
	if err != nil {
		log.Fatal(err)
	}

	var jsonMap JsonMap
	err = json.Unmarshal(jsonMapContents, &jsonMap)
	if err != nil {
		log.Fatal(err)
	}

	var jsonCollisions JsonSprites
	var jsonCollisionsGidStart int
	err = json.Unmarshal(jsonCollisionsContents, &jsonCollisions)
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

	for _, tileSet := range jsonMap.TileSets {
		switch tileSet.Source {
		case "collisions.tsj":
			jsonCollisionsGidStart = tileSet.FirstGid
		case "items.tsj":
			jsonItemsGidStart = tileSet.FirstGid
		case "bloons.tsj":
			jsonBloonsGidStart = tileSet.FirstGid
		}
	}

	itemsSprites := prepareSprites(itemSpritesRaw, &jsonItems, jsonItemsGidStart)
	collisionsSprites := prepareSprites(collisionSpritesRaw, &jsonCollisions, jsonCollisionsGidStart)
	bloonsSprites := prepareSprites(bloonsSpritesRaw, &jsonBloons, jsonBloonsGidStart)

	chiliAnimations := &Sprites{
		Texture:      chiliAnimationsRaw,
		TileWidth:    jsonItems.TileWidth,
		TileHeight:   jsonItems.TileHeight,
		WidthInTiles: int(chiliAnimationsRaw.Width) / jsonItems.TileWidth,
	}

	itemObjects := prepareObjects(jsonMap, itemsSprites, 1, jsonItemsGidStart, jsonItems.TileCount)
	collisionObjects := prepareObjects(jsonMap, collisionsSprites, 2, jsonCollisionsGidStart, jsonCollisions.TileCount)
	bloonsObjects := prepareBloons(jsonMap, bloonsSprites, 1, jsonBloonsGidStart, jsonBloons.TileCount)
	player := &Player{}

	player = &Player{
		IsMoving:       false,
		AnimationStep:  0,
		Rotation:       0,
		Animation:      chiliAnimations,
		DashLastUse:    time.Time{},
		DashCooldown:   time.Millisecond * 1200,
		AttackLastUse:  time.Time{},
		AttackCooldown: time.Millisecond * 500,
		Object: Object{
			Position:  rl.Vector2{X: 0 * global.Scale, Y: 0 * global.Scale},
			Texture:   chiliAnimations.Texture,
			Rectangle: chiliAnimations.GetSrc(7),
			HitBox: rl.Rectangle{
				X:      0 * global.Scale,
				Y:      0 * global.Scale,
				Width:  float32(global.TileWidth) * global.Scale,
				Height: float32(global.TileWidth) * global.Scale,
			},
		}}

	return itemObjects, collisionObjects, bloonsObjects, player
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
			}
		}
		sprites.Properties = append(sprites.Properties, property)
	}

	return sprites
}

func prepareObjects(jsonMap JsonMap, sprites *Sprites, layer, startId, n int) []*Object {
	objects := make([]*Object, 0)

	for i, val := range jsonMap.Layers[layer].Data {
		object := prepareObject(i, val, startId, n, sprites)
		if object == nil {
			continue
		}

		objects = append(objects, object)
	}

	return objects
}

func prepareBloons(jsonMap JsonMap, sprites *Sprites, layer, startId, n int) []*Bloon {
	bloons := make([]*Bloon, 0)

	for i, val := range jsonMap.Layers[layer].Data {
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

	x := float32(i % global.MapWidth * global.TileWidth)
	y := float32(i / global.MapWidth * global.TileWidth)

	object := &Object{
		Position:  rl.Vector2{X: x * global.Scale, Y: y * global.Scale},
		Texture:   sprites.Texture,
		Rectangle: sprites.GetSrc(val - startId),
	}

	for _, property := range sprites.Properties {
		if property.id == val {
			object.HitBox = rl.Rectangle{
				X:      (x + float32(property.HitBoxX)) * global.Scale,
				Y:      (y + float32(property.HitBoxY)) * global.Scale,
				Width:  float32(property.HitBoxWidth) * global.Scale,
				Height: float32(property.HitBoxHeight) * global.Scale,
			}
			break
		}
	}

	return object
}
