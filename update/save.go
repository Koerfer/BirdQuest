package update

import (
	"BirdQuest/global"
	"BirdQuest/menus"
	"BirdQuest/movement"
	"BirdQuest/save"
	"BirdQuest/scene"
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func LoadHandler(player *models.Player, camera *rl.Camera2D) (*models.Player, *rl.Camera2D) {
	saveState := save.Load()
	if saveState == nil {
		return nil, nil
	}
	global.UnloadAllTextures()

	player = saveState.Player
	camera = &saveState.Camera
	scene.CurrentScene = saveState.CurrentScene
	scene.AllScenes = saveState.Scenes
	global.VariableSet = saveState.GlobalVariables
	global.LoadAllTextures()

	rl.SetWindowSize(int(saveState.WindowWidth), int(saveState.WindowHeight))
	rl.SetWindowPosition(int(saveState.WindowPosition.X), int(saveState.WindowPosition.Y))
	updateDesiredWindowSize(saveState.WindowWidth, saveState.WindowHeight, player, camera)

	if saveState.IsFullScreen && !rl.IsWindowFullscreen() {
		rl.ToggleFullscreen()
	}

	return player, camera
}

func InitialLoader() (*models.Player, rl.Camera2D) {
	saveState := save.Load()
	if saveState == nil {
		global.SetDesiredWindowSize(1920, 1080)
		global.SetFPS(120)
		global.LoadAllTextures()

		camera := rl.Camera2D{}
		camera.Target = rl.Vector2{}
		global.Zoom(1, &camera)

		player := scene.SetScene("main", 250, 250, nil)
		movement.InitialiseCamera(player, &camera)

		return player, camera
	}

	global.SetDesiredWindowSize(saveState.WindowWidth, saveState.WindowHeight)
	global.SetFPS(120)

	camera := saveState.Camera
	player := saveState.Player

	global.VariableSet = saveState.GlobalVariables
	global.LoadAllTextures()

	for sceneName, savedScene := range saveState.Scenes {
		if sceneName == saveState.CurrentScene.Name {
			continue
		}

		scene.SetScene(sceneName, 250, 250, nil)
		scene.CurrentScene.CollisionObjects.DrawDynamic = savedScene.CollisionObjects.DrawDynamic
		scene.CurrentScene.CollisionObjects.DrawFirst = savedScene.CollisionObjects.DrawFirst
		scene.CurrentScene.CollisionObjects.DrawLast = savedScene.CollisionObjects.DrawLast
		scene.CurrentScene.BaseCollisionBoxes = savedScene.BaseCollisionBoxes
		scene.CurrentScene.CollisionBoxes = savedScene.CollisionBoxes
		scene.CurrentScene.Bloons.BloonObjects = savedScene.Bloons.BloonObjects
		scene.CurrentScene.Doors = savedScene.Doors
		scene.CurrentScene.ItemObjects.Objects = savedScene.ItemObjects.Objects
	}

	scene.SetScene(saveState.CurrentScene.Name, 250, 250, nil)
	scene.CurrentScene.CollisionObjects.DrawDynamic = saveState.CurrentScene.CollisionObjects.DrawDynamic
	scene.CurrentScene.CollisionObjects.DrawFirst = saveState.CurrentScene.CollisionObjects.DrawFirst
	scene.CurrentScene.CollisionObjects.DrawLast = saveState.CurrentScene.CollisionObjects.DrawLast
	scene.CurrentScene.BaseCollisionBoxes = saveState.CurrentScene.BaseCollisionBoxes
	scene.CurrentScene.CollisionBoxes = saveState.CurrentScene.CollisionBoxes
	scene.CurrentScene.Bloons.BloonObjects = saveState.CurrentScene.Bloons.BloonObjects
	scene.CurrentScene.Doors = saveState.CurrentScene.Doors
	scene.CurrentScene.ItemObjects.Objects = saveState.CurrentScene.ItemObjects.Objects

	menus.AllMenus = saveState.AllMenus

	rl.SetWindowPosition(int(saveState.WindowPosition.X), int(saveState.WindowPosition.Y))
	if saveState.IsMaximised {
		rl.MaximizeWindow()
	}

	if saveState.IsFullScreen {
		rl.ToggleFullscreen()
	}

	save.Save(player, camera)
	return player, camera
}
