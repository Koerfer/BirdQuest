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
	Bloons           *Bloons

	Width  float32
	Height float32

	WidthInTiles  int
	HeightInTiles int
}

var CurrentScene *Scene

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

	CurrentScene = &Scene{
		Name:          sceneName,
		Background:    background,
		Width:         float32(background.Width),
		WidthInTiles:  int(background.Width) / global.TileWidth,
		Height:        float32(background.Height),
		HeightInTiles: int(background.Height / global.TileHeight),
	}

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
	initiateObjects(scenePath)

	player := preparePlayer(playerPath)
	player.Position.X = playerX * global.VariableSet.EntityScale
	player.Position.Y = playerY * global.VariableSet.EntityScale
	CurrentScene.Background = background

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
