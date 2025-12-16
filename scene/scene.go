package scene

import (
	"BirdQuest/global"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
	"os"
	"path/filepath"
)

type Scene struct {
	Name               string
	Background         rl.Texture2D
	ItemObjects        []*Object
	CollisionObjects   []*rl.Rectangle
	CollisionObjects3d []*Object
	BloonObjects       []*Bloon

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
	path := filepath.Join(cwd, "sprites", sceneName)
	background := rl.LoadTexture(filepath.Join(path, "background.png"))
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

	itemObjects, collisionObjects, collisionObjects3d, bloonObjects, player := InitiateObjects(path)

	player.Position.X = playerX * global.VariableSet.EntityScale
	player.Position.Y = playerY * global.VariableSet.EntityScale
	CurrentScene.Background = background
	CurrentScene.ItemObjects = itemObjects
	CurrentScene.CollisionObjects = collisionObjects
	CurrentScene.CollisionObjects3d = collisionObjects3d
	CurrentScene.BloonObjects = bloonObjects

	return player
}

func UnloadAllTextures() {
	rl.UnloadTexture(CurrentScene.Background)

	if CurrentScene.ItemObjects != nil && len(CurrentScene.ItemObjects) != 0 {
		rl.UnloadTexture(CurrentScene.ItemObjects[0].Texture)
	}
	if CurrentScene.CollisionObjects3d != nil && len(CurrentScene.CollisionObjects3d) != 0 {
		rl.UnloadTexture(CurrentScene.CollisionObjects3d[0].Texture)
	}
	if CurrentScene.BloonObjects != nil && len(CurrentScene.BloonObjects) != 0 {
		rl.UnloadTexture(CurrentScene.BloonObjects[0].Texture)
	}

	CurrentScene = nil
}
