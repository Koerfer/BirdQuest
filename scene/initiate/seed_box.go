package initiate

import (
	"BirdQuest/global"
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func prepareSeedBoxes(jMap *jsonMap, layerName string, startId int, scene *models.Scene, jSprites *jsonSprites) {
	seedBoxSprites := prepareSprites(global.VariableSet.Textures32x32, jSprites, startId)

	scene.SeedBoxes = make([]*models.SeedBox, 0)

	for _, layer := range jMap.Layers {
		if layer.Name != layerName {
			continue
		}

		for i, val := range layer.Data {
			prepareSeedBox(i, val, startId, jSprites.TileCount, seedBoxSprites, scene)
		}
	}
}

func prepareSeedBox(i, val, startId, n int, sprites *models.Sprites, scene *models.Scene) {
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

	scene.SeedBoxes = append(scene.SeedBoxes, &models.SeedBox{
		OpeningStage: 0,
		Object:       *object,
	})
}
