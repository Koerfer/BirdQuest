package initiate

import (
	"BirdQuest/global"
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func prepareCollisions(jsonMap *jsonMap, scene *models.Scene) {
	baseCollisionBoxes := make([]*rl.Rectangle, 0)

	for _, jsonMapLayer := range jsonMap.Layers {
		if jsonMapLayer.Name != "Collisions" {
			continue
		}

		for _, obstacle := range jsonMapLayer.Objects {
			baseCollisionBoxes = append(baseCollisionBoxes, &rl.Rectangle{
				X:      obstacle.X,
				Y:      obstacle.Y,
				Width:  obstacle.Width,
				Height: obstacle.Height,
			})
		}
	}

	scene.BaseCollisionBoxes = baseCollisionBoxes
}

func prepareCollisionObjects(jMap *jsonMap, layerName string, startId int, scene *models.Scene, jSprites *jsonSprites) {
	collisionSprites := prepareSprites(global.VariableSet.Textures32x32, jSprites, startId)

	scene.CollisionObjects = &models.CollisionItems{
		DrawFirst:   make([]*models.Object, 0),
		DrawDynamic: make([]*models.Object, 0),
		DrawLast:    make([]*models.Object, 0),
	}

	for _, layer := range jMap.Layers {
		if layer.Name != layerName {
			continue
		}

		for i, val := range layer.Data {
			prepareCollisionObject(i, val, startId, jSprites.TileCount, collisionSprites, scene)
		}
	}
}

func prepareCollisionObject(i, val, startId, n int, sprites *models.Sprites, scene *models.Scene) {
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
