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

func prepareCollisions(jsonMap *jsonMap, scene *models.Scene) {
	baseCollisionBoxes := make([]*rl.Rectangle, 0)
	collisionBoxes := make([]*rl.Rectangle, 0)

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
			collisionBoxes = append(collisionBoxes, &rl.Rectangle{
				X:      obstacle.X * global.VariableSet.EntityScale,
				Y:      obstacle.Y * global.VariableSet.EntityScale,
				Width:  obstacle.Width * global.VariableSet.EntityScale,
				Height: obstacle.Height * global.VariableSet.EntityScale,
			})
		}
	}

	scene.BaseCollisionBoxes = baseCollisionBoxes
	scene.CollisionBoxes = collisionBoxes
}

func prepareCollisionObjects(jMap *jsonMap, layerName, path string, startId int, scene *models.Scene) {
	jsonCollisionContents, err := os.ReadFile(filepath.Join(path, "collisions.tsj"))
	if err != nil {
		log.Fatal(err)
	}

	var jsonCollisions jsonSprites
	err = json.Unmarshal(jsonCollisionContents, &jsonCollisions)
	if err != nil {
		log.Fatal(err)
	}

	scene.CollisionObjects = &models.CollisionItems{
		DrawFirst:   make([]*models.Object, 0),
		DrawDynamic: make([]*models.Object, 0),
		DrawLast:    make([]*models.Object, 0),
	}

	collisionSprites := prepareSprites(global.VariableSet.CollisionObjectsTexture, &jsonCollisions, startId)

	for _, layer := range jMap.Layers {
		if layer.Name != layerName {
			continue
		}

		for i, val := range layer.Data {
			prepareCollisionObject(i, val, startId, jsonCollisions.TileCount, collisionSprites, scene)
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
		BasePosition:  &rl.Vector2{X: x, Y: y},
		BaseRectangle: sprites.GetRectangleAreaInTexture(val - startId),
	}
	object.Rectangle = &rl.Rectangle{
		X:      object.BasePosition.X * global.VariableSet.EntityScale,
		Y:      object.BasePosition.Y * global.VariableSet.EntityScale,
		Width:  object.BaseRectangle.Width * global.VariableSet.EntityScale,
		Height: object.BaseRectangle.Height * global.VariableSet.EntityScale,
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
