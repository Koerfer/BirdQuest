package movement

import (
	"BirdQuest/global"
	"BirdQuest/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var screenHeight float32 = global.ScreenHeight * global.Scale
var screenWidth float32 = global.ScreenWidth * global.Scale
var playerMiddleOffset float32 = 16 * global.Scale

func CorrectForZoom(x, y float32, camera *rl.Camera2D) {
	if camera.Zoom == 1 {
		camera.Target.X = 0
		camera.Target.Y = 0
		return
	}

	if x < 0 {
		camera.Target.X = 0
	} else if x+screenWidth/camera.Zoom >= screenWidth {
		camera.Target.X = screenWidth - screenWidth/camera.Zoom
	} else {
		camera.Target.X = x
	}

	if y < 0 {
		camera.Target.Y = 0
	} else if y+screenHeight/camera.Zoom >= screenHeight {
		camera.Target.Y = screenHeight - screenHeight/camera.Zoom
	} else {
		camera.Target.Y = y
	}
}

func InitialiseCamera(player *objects.Player, camera *rl.Camera2D, offsetX, offsetY float32) {
	moveCameraRight(player, camera, offsetX)
	moveCameraDown(player, camera, offsetY)
}

func moveCameraUp(player *objects.Player, camera *rl.Camera2D, offset float32) {
	if player.Position.Y+playerMiddleOffset <= camera.Target.Y+screenHeight/(camera.Zoom*2) && camera.Target.Y > offset {
		camera.Target.Y -= offset
	} else if camera.Target.Y <= offset {
		camera.Target.Y = 0
	}
}

func moveCameraDown(player *objects.Player, camera *rl.Camera2D, offset float32) {
	stopFactor := (camera.Zoom - 1) / camera.Zoom

	if player.Position.Y+playerMiddleOffset >= camera.Target.Y+screenHeight/(camera.Zoom*2) && camera.Target.Y+offset < screenHeight*stopFactor {
		camera.Target.Y += offset
	} else if camera.Target.Y+offset >= screenHeight*stopFactor {
		camera.Target.Y = screenHeight * stopFactor
	}
}

func moveCameraLeft(player *objects.Player, camera *rl.Camera2D, offset float32) {
	if player.Position.X+playerMiddleOffset <= camera.Target.X+screenWidth/(camera.Zoom*2) && camera.Target.X-offset > 0 {
		camera.Target.X -= offset
	} else if camera.Target.X-offset <= 0 {
		camera.Target.X = 0
	}
}

func moveCameraRight(player *objects.Player, camera *rl.Camera2D, offset float32) {
	stopFactor := (camera.Zoom - 1) / camera.Zoom

	if player.Position.X+playerMiddleOffset >= camera.Target.X+screenWidth/(camera.Zoom*2) && camera.Target.X+offset < screenWidth*stopFactor {
		camera.Target.X += offset
	} else if camera.Target.X+offset >= screenWidth*stopFactor {
		camera.Target.X = screenWidth * stopFactor
	}
}
