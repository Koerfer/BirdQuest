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
	scene.AllScenes[scene.CurrentScene.Name] = scene.CurrentScene
	scene.Quests = saveState.Quests
	scene.Quests[0].Steps[1].Box = scene.AllScenes["main"].SeedBoxes[0]
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

		scene.CreateQuests()

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
		scene.AllScenes[sceneName].CollisionObjects.DrawDynamic = savedScene.CollisionObjects.DrawDynamic
		scene.AllScenes[sceneName].CollisionObjects.DrawFirst = savedScene.CollisionObjects.DrawFirst
		scene.AllScenes[sceneName].CollisionObjects.DrawLast = savedScene.CollisionObjects.DrawLast
		scene.AllScenes[sceneName].BaseCollisionBoxes = savedScene.BaseCollisionBoxes
		scene.AllScenes[sceneName].Bloons.BloonObjects = savedScene.Bloons.BloonObjects
		scene.AllScenes[sceneName].Doors = savedScene.Doors
		scene.AllScenes[sceneName].ItemObjects.Objects = savedScene.ItemObjects.Objects
		scene.AllScenes[sceneName].NPCs = savedScene.NPCs
		scene.AllScenes[sceneName].SeedBoxes = savedScene.SeedBoxes
	}

	scene.SetScene(saveState.CurrentScene.Name, 250, 250, nil)
	scene.CurrentScene.CollisionObjects.DrawDynamic = saveState.CurrentScene.CollisionObjects.DrawDynamic
	scene.CurrentScene.CollisionObjects.DrawFirst = saveState.CurrentScene.CollisionObjects.DrawFirst
	scene.CurrentScene.CollisionObjects.DrawLast = saveState.CurrentScene.CollisionObjects.DrawLast
	scene.CurrentScene.BaseCollisionBoxes = saveState.CurrentScene.BaseCollisionBoxes
	scene.CurrentScene.Bloons.BloonObjects = saveState.CurrentScene.Bloons.BloonObjects
	scene.CurrentScene.Doors = saveState.CurrentScene.Doors
	scene.CurrentScene.ItemObjects.Objects = saveState.CurrentScene.ItemObjects.Objects
	scene.CurrentScene.NPCs = saveState.CurrentScene.NPCs
	scene.CurrentScene.SeedBoxes = saveState.CurrentScene.SeedBoxes

	menus.AllMenus = saveState.AllMenus
	scene.Quests = saveState.Quests
	scene.Quests[0].Steps[1].Box = scene.AllScenes["main"].SeedBoxes[0]

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
