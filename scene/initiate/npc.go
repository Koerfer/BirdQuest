package initiate

import (
	"BirdQuest/global"
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func prepareNPCs(jMap *jsonMap, layerName string, startId int, scene *models.Scene, jSprites *jsonSprites) {
	npcSprites := prepareSprites(global.VariableSet.Textures32x32, jSprites, startId)

	scene.NPCs = make([]*models.NPC, 0)

	for _, layer := range jMap.Layers {
		if layer.Name != layerName {
			continue
		}

		for i, val := range layer.Data {
			prepareNPC(i, val, startId, jSprites.TileCount, npcSprites, scene)
		}
	}
}

func prepareNPC(i, val, startId, n int, sprites *models.Sprites, scene *models.Scene) {
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

	for _, prop := range sprites.Properties {
		if prop.Id == val {
			scene.NPCs = append(scene.NPCs, &models.NPC{
				Name:          prop.Name,
				StartedQuests: make([]int, 0),
				Object:        *object,
			})
			return
		}
	}
}
