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
				X:      obstacle.X * global.Scale,
				Y:      obstacle.Y * global.Scale,
				Width:  obstacle.Width * global.Scale,
				Height: obstacle.Height * global.Scale,
			})
		}
	}

	return collisionObjects
}
