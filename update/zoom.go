package update

import (
	"BirdQuest/global"
	"BirdQuest/movement"
	"BirdQuest/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var minZoom float32 = 1
var maxZoom float32 = 10

func updateZoom(camera *rl.Camera2D, player *objects.Player) {
	if rl.GetMouseWheelMove() != 0 {
		camera.Zoom += rl.GetMouseWheelMove() / 10 * camera.Zoom
		if camera.Zoom < minZoom {
			camera.Zoom = minZoom
		} else if camera.Zoom > maxZoom {
			camera.Zoom = maxZoom
		}
		newX := player.Position.X + 16*global.Scale - global.ScreenWidth*global.Scale/camera.Zoom/2
		newY := player.Position.Y + 16*global.Scale - global.ScreenHeight*global.Scale/camera.Zoom/2

		movement.CorrectForZoom(newX, newY, camera)
	}
}
