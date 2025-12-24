package update

import (
	"BirdQuest/global"
	"BirdQuest/movement"
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var minZoom float32 = 1
var maxZoom float32 = 10

func updateZoom(camera *rl.Camera2D, player *models.Player) {
	if rl.GetMouseWheelMove() != 0 {
		camera.Zoom += rl.GetMouseWheelMove() / 10 * camera.Zoom
		if camera.Zoom < minZoom {
			camera.Zoom = minZoom
		} else if camera.Zoom > maxZoom {
			camera.Zoom = maxZoom
		}
		global.Zoom(camera.Zoom, camera)

		newX := player.Rectangle.X + global.VariableSet.PlayerMiddleOffset - global.VariableSet.VisibleMapWidth/2
		newY := player.Rectangle.Y + global.VariableSet.PlayerMiddleOffset - global.VariableSet.VisibleMapHeight/2

		movement.CorrectForZoom(newX, newY, camera)
	}
}
