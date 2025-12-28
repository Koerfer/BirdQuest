package initiate

import (
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func prepareDoors(jsonMap *jsonMap, scene *models.Scene) {
	doors := make([]*models.Door, 0)

	for _, jsonMapLayer := range jsonMap.Layers {
		if jsonMapLayer.Name != "Doors" {
			continue
		}

		for _, jsonDoor := range jsonMapLayer.Objects {
			door := &models.Door{
				BaseRectangle: &rl.Rectangle{
					X:      jsonDoor.X,
					Y:      jsonDoor.Y,
					Width:  jsonDoor.Width,
					Height: jsonDoor.Height,
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
