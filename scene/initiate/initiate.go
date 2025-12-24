package initiate

import (
	"BirdQuest/scene/models"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type jsonTiles struct {
	Id         int             `json:"id"`
	Properties []*jsonProperty `json:"properties"`
}

func Objects(path string, scene *models.Scene) {
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
