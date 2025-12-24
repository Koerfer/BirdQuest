package update

import (
	"BirdQuest/global"
	"BirdQuest/save"
	"BirdQuest/scene"
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func SaveHandler(player *models.Player, camera rl.Camera2D) {
	if rl.IsKeyPressed(rl.KeyF6) {
		save.Save(player, camera)
	}

	if rl.IsKeyPressed(rl.KeyF7) {
		saveState := save.Load()
		player = saveState.Player
		camera = saveState.Camera
		scene.CurrentScene = saveState.CurrentScene
		scene.AllScenes = saveState.Scenes
		global.VariableSet = saveState.GlobalVariables
	}
}
