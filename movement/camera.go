package movement

import (
	"BirdQuest/global"
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func CorrectForZoom(x, y float32, camera *rl.Camera2D) {
	if x < 0 {
		camera.Target.X = 0
	} else if x+global.VariableSet.VisibleMapWidth >= global.VariableSet.MapWidth {
		camera.Target.X = global.VariableSet.MapWidth - global.VariableSet.VisibleMapWidth
	} else {
		camera.Target.X = x
	}

	if y < 0 {
		camera.Target.Y = 0
	} else if y+global.VariableSet.VisibleMapHeight >= global.VariableSet.MapHeight {
		camera.Target.Y = global.VariableSet.MapHeight - global.VariableSet.VisibleMapHeight
	} else {
		camera.Target.Y = y
	}
}

func InitialiseCamera(player *models.Player, camera *rl.Camera2D) {
	newX := player.Rectangle.X + global.VariableSet.PlayerMiddleOffset - global.VariableSet.VisibleMapWidth/2
	newY := player.Rectangle.Y + global.VariableSet.PlayerMiddleOffset - global.VariableSet.VisibleMapHeight/2

	CorrectForZoom(newX, newY, camera)
}

func moveCameraUp(player *models.Player, camera *rl.Camera2D, offset float32) {
	if player.Rectangle.Y+global.VariableSet.PlayerMiddleOffset <= camera.Target.Y+global.VariableSet.DesiredHeight/(camera.Zoom*2) &&
		camera.Target.Y > offset {

		camera.Target.Y -= offset
	} else if camera.Target.Y <= offset {
		camera.Target.Y = 0
	}
}

func moveCameraDown(player *models.Player, camera *rl.Camera2D, offset float32) {
	if player.Rectangle.Y+global.VariableSet.PlayerMiddleOffset >= camera.Target.Y+global.VariableSet.VisibleMapHeight/2 &&
		camera.Target.Y+global.VariableSet.VisibleMapHeight+offset < global.VariableSet.MapHeight {

		camera.Target.Y += offset
	} else if camera.Target.Y+global.VariableSet.VisibleMapHeight+offset >= global.VariableSet.MapHeight {
		camera.Target.Y = global.VariableSet.MapHeight - global.VariableSet.VisibleMapHeight
	}
}

func moveCameraLeft(player *models.Player, camera *rl.Camera2D, offset float32) {
	if player.Rectangle.X+global.VariableSet.PlayerMiddleOffset <= camera.Target.X+global.VariableSet.DesiredWidth/(camera.Zoom*2) &&
		camera.Target.X-offset > 0 {

		camera.Target.X -= offset
	} else if camera.Target.X-offset <= 0 {
		camera.Target.X = 0
	}
}

func moveCameraRight(player *models.Player, camera *rl.Camera2D, offset float32) {
	if player.Rectangle.X+global.VariableSet.PlayerMiddleOffset >= camera.Target.X+global.VariableSet.VisibleMapWidth/2 &&
		camera.Target.X+global.VariableSet.VisibleMapWidth+offset < global.VariableSet.MapWidth {

		camera.Target.X += offset
	} else if camera.Target.X+global.VariableSet.VisibleMapWidth+offset >= global.VariableSet.MapWidth {
		camera.Target.X = global.VariableSet.MapWidth - global.VariableSet.VisibleMapWidth
	}
}
