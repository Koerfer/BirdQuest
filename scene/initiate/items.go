package initiate

import (
	"BirdQuest/global"
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func prepareItems(jMap *jsonMap, layerName string, startId int, scene *models.Scene, jSprites *jsonSprites) {
	scene.ItemObjects = &models.Items{
		Objects: make([]*models.Object, 0),
	}

	itemSprites := prepareSprites(global.VariableSet.Textures32x32, jSprites, startId)

	for _, layer := range jMap.Layers {
		if layer.Name != layerName {
			continue
		}

		for i, val := range layer.Data {
			prepareItem(i, val, startId, jSprites.TileCount, itemSprites, scene)
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
