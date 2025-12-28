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

func prepareItems(jMap *jsonMap, layerName, path string, startId int, scene *models.Scene) {
	jsonItemContents, err := os.ReadFile(filepath.Join(path, "items.tsj"))
	if err != nil {
		log.Fatal(err)
	}
	var jsonItems jsonSprites
	err = json.Unmarshal(jsonItemContents, &jsonItems)
	if err != nil {
		log.Fatal(err)
	}

	scene.ItemObjects = &models.Items{
		Objects: make([]*models.Object, 0),
	}

	itemSprites := prepareSprites(global.VariableSet.ItemsTexture, &jsonItems, startId)

	for _, layer := range jMap.Layers {
		if layer.Name != layerName {
			continue
		}

		for i, val := range layer.Data {
			prepareItem(i, val, startId, jsonItems.TileCount, itemSprites, scene)
		}
	}
}

func prepareItem(i, val, startId, n int, sprites *models.Sprites, scene *models.Scene) {
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

	object := &models.Object{
		BaseRectangle: sprites.GetRectangleAreaInTexture(val - startId),
	}
	object.BasePositionRectangle = &rl.Rectangle{
		X:      x,
		Y:      y,
		Width:  object.BaseRectangle.Width,
		Height: object.BaseRectangle.Height,
	}

	scene.ItemObjects.Objects = append(scene.ItemObjects.Objects, object)
}
