package scene

import (
	"BirdQuest/global"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func prepareCollisions(jsonMap *jsonMap, scene *Scene) {
	collisionBoxes := make([]*rl.Rectangle, 0)

	for _, jsonMapLayer := range jsonMap.Layers {
		if jsonMapLayer.Name != "Collisions" {
			continue
		}

		for _, obstacle := range jsonMapLayer.Objects {
			collisionBoxes = append(collisionBoxes, &rl.Rectangle{
				X:      obstacle.X * global.VariableSet.EntityScale,
				Y:      obstacle.Y * global.VariableSet.EntityScale,
				Width:  obstacle.Width * global.VariableSet.EntityScale,
				Height: obstacle.Height * global.VariableSet.EntityScale,
			})
		}
	}

	scene.CollisionBoxes = collisionBoxes
}

func prepareDoors(jsonMap *jsonMap, scene *Scene) {
	doors := make([]*Door, 0)

	for _, jsonMapLayer := range jsonMap.Layers {
		if jsonMapLayer.Name != "Doors" {
			continue
		}

		for _, jsonDoor := range jsonMapLayer.Objects {
			door := &Door{
				Rectangle: &rl.Rectangle{
					X:      jsonDoor.X * global.VariableSet.EntityScale,
					Y:      jsonDoor.Y * global.VariableSet.EntityScale,
					Width:  jsonDoor.Width * global.VariableSet.EntityScale,
					Height: jsonDoor.Height * global.VariableSet.EntityScale,
				},
			}

			for _, prop := range jsonDoor.Properties {
				switch prop.Name {
				case "goesTo":
					door.GoesToScene = prop.Value.(string)
				case "goesToX":
					door.GoesToX = float32(prop.Value.(float64))
				case "goesToY":
					door.GoesToY = float32(prop.Value.(float64))
				}
			}

			doors = append(doors, door)
		}
	}

	scene.Doors = doors
}
