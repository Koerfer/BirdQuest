package scene

import (
	"BirdQuest/global"
	"BirdQuest/scene/initiate"
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
	"os"
	"path/filepath"
)

var CurrentScene *models.Scene
var AllScenes map[string]*models.Scene

func ChangeScene(door *models.Door, player *models.Player) *models.Player {
	if AllScenes[door.GoesToScene] == nil {
		return SetScene(door.GoesToScene, door.GoesToX, door.GoesToY, player)
	}

	player.DashDirection.X *= 1 / global.VariableSet.Speed
	player.DashDirection.Y *= 1 / global.VariableSet.Speed

	CurrentScene = AllScenes[door.GoesToScene]

	global.VariableSet.EntityScale = global.VariableSet.DesiredWidth / CurrentScene.Width
	global.VariableSet.ScaleHeight = global.VariableSet.DesiredHeight / CurrentScene.Height
	global.VariableSet.ScaleWidth = global.VariableSet.DesiredWidth / CurrentScene.Width

	if global.VariableSet.EntityScale < global.VariableSet.DesiredHeight/CurrentScene.Height {
		global.VariableSet.EntityScale = global.VariableSet.DesiredHeight / CurrentScene.Height
	}

	global.VariableSet.MapHeight = CurrentScene.Height * global.VariableSet.EntityScale
	global.VariableSet.MapWidth = CurrentScene.Width * global.VariableSet.EntityScale

	global.VariableSet.BasePlayerMiddleOffset = global.TileWidth / 2
	global.VariableSet.PlayerMiddleOffset = global.TileWidth / 2 * global.VariableSet.EntityScale
	global.VariableSet.EntitySize = global.TileWidth * global.VariableSet.EntityScale

	global.VariableSet.Speed = global.VariableSet.FpsScale * global.VariableSet.EntityScale

	player.Rectangle.Width = player.BaseRectangle.Width * global.VariableSet.EntityScale
	player.Rectangle.Height = player.BaseRectangle.Height * global.VariableSet.EntityScale

	player.BasePosition.X = door.GoesToX
	player.BasePosition.Y = door.GoesToY

	player.Rectangle.X = door.GoesToX * global.VariableSet.EntityScale
	player.Rectangle.Y = door.GoesToY * global.VariableSet.EntityScale

	player.DashDirection.X *= global.VariableSet.Speed
	player.DashDirection.Y *= global.VariableSet.Speed

	return player
}

func SetScene(sceneName string, playerX, playerY float32, player *models.Player) *models.Player {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	scenePath := filepath.Join(cwd, "sprites", sceneName)
	background := rl.LoadTexture(filepath.Join(scenePath, "background.png"))
	if background.Width == 0 {
		return SetScene(sceneName, playerX, playerY, player)
	}

	scene := &models.Scene{
		Name:          sceneName,
		Background:    background,
		Width:         float32(background.Width),
		WidthInTiles:  int(background.Width) / global.TileWidth,
		Height:        float32(background.Height),
		HeightInTiles: int(background.Height / global.TileHeight),
	}

	global.VariableSet.EntityScale = global.VariableSet.DesiredWidth / scene.Width
	global.VariableSet.ScaleHeight = global.VariableSet.DesiredHeight / scene.Height
	global.VariableSet.ScaleWidth = global.VariableSet.DesiredWidth / scene.Width

	if global.VariableSet.EntityScale < global.VariableSet.DesiredHeight/scene.Height {
		global.VariableSet.EntityScale = global.VariableSet.DesiredHeight / scene.Height
	}

	global.VariableSet.MapHeight = scene.Height * global.VariableSet.EntityScale
	global.VariableSet.MapWidth = scene.Width * global.VariableSet.EntityScale

	global.VariableSet.BasePlayerMiddleOffset = global.TileWidth / 2
	global.VariableSet.PlayerMiddleOffset = global.TileWidth / 2 * global.VariableSet.EntityScale
	global.VariableSet.EntitySize = global.TileWidth * global.VariableSet.EntityScale

	global.VariableSet.Speed = global.VariableSet.FpsScale * global.VariableSet.EntityScale
	initiate.InitiateObjects(scenePath, scene)

	if player == nil {
		player = initiate.PreparePlayer()
	} else {
		player.Rectangle.Width = player.BaseRectangle.Width * global.VariableSet.EntityScale
		player.Rectangle.Height = player.BaseRectangle.Height * global.VariableSet.EntityScale
	}

	player.BasePosition.X = playerX
	player.BasePosition.Y = playerY

	player.Rectangle.X = playerX * global.VariableSet.EntityScale
	player.Rectangle.Y = playerY * global.VariableSet.EntityScale

	scene.Background = background

	if AllScenes == nil {
		AllScenes = make(map[string]*models.Scene)
	}
	AllScenes[sceneName] = scene
	CurrentScene = scene

	return player
}

func UnloadAllBackgroundTextures() {
	for _, scene := range AllScenes {
		rl.UnloadTexture(scene.Background)
	}
}
