package initiate

import (
	"BirdQuest/global"
	"BirdQuest/scene/models"
	"encoding/json"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
	"os"
	"path/filepath"
)

func prepareBloons(jMap *jsonMap, layerName, path string, startId int, scene *models.Scene) {
	jsonBloonContents, err := os.ReadFile(filepath.Join(path, "bloons.tsj"))
	if err != nil {
		log.Fatal(err)
	}
	var jsonBloons jsonSprites
	err = json.Unmarshal(jsonBloonContents, &jsonBloons)
	if err != nil {
		log.Fatal(err)
	}

	scene.Bloons = &models.Bloons{
		BloonObjects: make([]*models.Bloon, 0),
	}

	bloonSprites := prepareSprites(global.VariableSet.BloonsTexture, &jsonBloons, startId)

	for _, layer := range jMap.Layers {
		if layer.Name != layerName {
			continue
		}

		for i, val := range layer.Data {
			prepareBloon(i, val, startId, jsonBloons.TileCount, bloonSprites, scene)
		}
	}
}

func prepareBloon(i, val, startId, n int, sprites *models.Sprites, scene *models.Scene) {
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

	bloonObject := &models.Object{BaseRectangle: sprites.GetRectangleAreaInTexture(val - startId)}
	bloonObject.BasePositionRectangle = &rl.Rectangle{
		X:      x,
		Y:      y,
		Width:  bloonObject.BaseRectangle.Width,
		Height: bloonObject.BaseRectangle.Height,
	}

	for _, prop := range sprites.Properties {
		if prop.Id == val {
			bloon := &models.Bloon{
				Lives: startId + 3 - val,
			}
			bloon.Object = *bloonObject

			scene.Bloons.BloonObjects = append(scene.Bloons.BloonObjects, bloon)
			return
		}
	}
}
