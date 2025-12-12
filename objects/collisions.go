package objects

import (
	"BirdQuest/global"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func prepareCollisions(jsonMap *JsonMap) []*rl.Rectangle {
	collisionObjects := make([]*rl.Rectangle, 0)

	for _, jsonMapLayer := range jsonMap.Layers {
		if jsonMapLayer.Name != "Collisions" {
			continue
		}

		for _, obstacle := range jsonMapLayer.Objects {
			collisionObjects = append(collisionObjects, &rl.Rectangle{
				X:      obstacle.X * global.VariableSet.EntityScale,
				Y:      obstacle.Y * global.VariableSet.EntityScale,
				Width:  obstacle.Width * global.VariableSet.EntityScale,
				Height: obstacle.Height * global.VariableSet.EntityScale,
			})
		}
	}

	return collisionObjects
}
