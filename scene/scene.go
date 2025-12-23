package scene

import (
	"BirdQuest/global"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
	"os"
	"path/filepath"
)

type Scene struct {
	Name             string
	Background       rl.Texture2D
	ItemObjects      *Items
	CollisionBoxes   []*rl.Rectangle
	CollisionObjects *CollisionItems
	Doors            []*Door
	Bloons           *Bloons

	Width  float32
	Height float32

	WidthInTiles  int
	HeightInTiles int
}

var CurrentScene *Scene
var AllScenes map[string]*Scene

func ChangeScene(door *Door, player *Player) *Player {
	if AllScenes[door.GoesToScene] == nil {
		return SetScene(door.GoesToScene, door.GoesToX, door.GoesToY)
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

	global.VariableSet.PlayerMiddleOffset = global.TileWidth / 2 * global.VariableSet.EntityScale
	global.VariableSet.EntitySize = global.TileWidth * global.VariableSet.EntityScale

	global.VariableSet.Speed = global.VariableSet.FpsScale * global.VariableSet.EntityScale

	player.Position.X = door.GoesToX * global.VariableSet.EntityScale
	player.Position.Y = door.GoesToY * global.VariableSet.EntityScale

	player.DashDirection.X *= global.VariableSet.Speed
	player.DashDirection.Y *= global.VariableSet.Speed

	return player
}

func SetScene(sceneName string, playerX, playerY float32) *Player {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	scenePath := filepath.Join(cwd, "sprites", sceneName)
	playerPath := filepath.Join(cwd, "sprites", "player")
	background := rl.LoadTexture(filepath.Join(scenePath, "background.png"))
	if background.Width == 0 {
		return SetScene(sceneName, playerX, playerY)
	}

	scene := &Scene{
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

	global.VariableSet.PlayerMiddleOffset = global.TileWidth / 2 * global.VariableSet.EntityScale
	global.VariableSet.EntitySize = global.TileWidth * global.VariableSet.EntityScale

	global.VariableSet.Speed = global.VariableSet.FpsScale * global.VariableSet.EntityScale
	initiateObjects(scenePath, scene)

	player := preparePlayer(playerPath)
	player.Position.X = playerX * global.VariableSet.EntityScale
	player.Position.Y = playerY * global.VariableSet.EntityScale
	scene.Background = background

	if AllScenes == nil {
		AllScenes = make(map[string]*Scene)
	}
	AllScenes[sceneName] = scene
	CurrentScene = scene

	return player
}

func UnloadAllTextures() {
	rl.UnloadTexture(CurrentScene.Background)

	if CurrentScene.ItemObjects != nil {
		rl.UnloadTexture(CurrentScene.ItemObjects.Texture)
	}
	if CurrentScene.CollisionObjects != nil {
		rl.UnloadTexture(CurrentScene.CollisionObjects.Texture)
	}
	if CurrentScene.Bloons != nil {
		rl.UnloadTexture(CurrentScene.Bloons.Texture)
	}

	CurrentScene = nil
}
