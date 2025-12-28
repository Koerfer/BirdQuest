package update

import (
	"BirdQuest/global"
	"BirdQuest/menus"
	"BirdQuest/movement"
	"BirdQuest/scene"
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Window(player *models.Player, camera *rl.Camera2D) {
	if rl.IsWindowResized() {
		updateDesiredWindowSize(float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight()), player, camera)
	}
}

func updateDesiredWindowSize(width, height float32, player *models.Player, camera *rl.Camera2D) {
	global.VariableSet.DesiredWidth = width
	global.VariableSet.DesiredHeight = height

	global.VariableSet.ScaleHeight = global.VariableSet.DesiredHeight / scene.CurrentScene.Height
	global.VariableSet.ScaleWidth = global.VariableSet.DesiredWidth / scene.CurrentScene.Width

	global.VariableSet.EntityScale = global.VariableSet.DesiredWidth / scene.CurrentScene.Width
	if global.VariableSet.EntityScale < global.VariableSet.DesiredHeight/scene.CurrentScene.Height {
		global.VariableSet.EntityScale = global.VariableSet.DesiredHeight / scene.CurrentScene.Height
	}

	global.VariableSet.MapHeight = scene.CurrentScene.Height * global.VariableSet.EntityScale
	global.VariableSet.MapWidth = scene.CurrentScene.Width * global.VariableSet.EntityScale

	global.VariableSet.PlayerMiddleOffset = global.TileWidth / 2 * global.VariableSet.EntityScale
	global.VariableSet.EntitySize = global.TileWidth * global.VariableSet.EntityScale

	global.VariableSet.VisibleMapHeight = global.VariableSet.DesiredHeight / camera.Zoom
	global.VariableSet.VisibleMapWidth = global.VariableSet.DesiredWidth / camera.Zoom

	for _, menu := range menus.AllMenus {
		menuScaler := global.VariableSet.EntityScale
		if menu.BaseRectangle.Height*global.VariableSet.EntityScale > global.VariableSet.VisibleMapHeight {
			menuScaler = global.VariableSet.VisibleMapHeight / menu.BaseRectangle.Height * 0.98 * camera.Zoom
		}

		menu.FontSize = menu.BaseFontSize * menuScaler
		menu.Rectangle.X = menu.BaseRectangle.X * menuScaler
		menu.Rectangle.Y = menu.BaseRectangle.Y * menuScaler
		menu.Rectangle.Width = menu.BaseRectangle.Width * menuScaler
		menu.Rectangle.Height = menu.BaseRectangle.Height * menuScaler

		for _, button := range menu.Buttons {
			button.Rectangle.X = button.BaseRectangle.X * menuScaler
			button.Rectangle.Y = button.BaseRectangle.Y * menuScaler
			button.Rectangle.Width = button.BaseRectangle.Width * menuScaler
			button.Rectangle.Height = button.BaseRectangle.Height * menuScaler
		}
	}

	movement.InitialiseCamera(player, camera)
}
