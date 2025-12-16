package scene

import (
	"BirdQuest/global"
	"BirdQuest/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
	"os"
	"path/filepath"
)

type Scene struct {
	Background         rl.Texture2D
	ItemObjects        []*objects.Object
	CollisionObjects   []*rl.Rectangle
	CollisionObjects3d []*objects.Object
	BloonObjects       []*objects.Bloon

	Width  float32
	Height float32
}

var CurrentScene *Scene

func SetScene(sceneName string, playerX, playerY float32) *objects.Player {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Join(cwd, "sprites", sceneName)
	background := rl.LoadTexture(filepath.Join(path, "Background.png"))
	scene := &Scene{
		Background: background,
		Width:      float32(background.Width),
		Height:     float32(background.Height),
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

	itemObjects, collisionObjects, collisionObjects3d, bloonObjects, player := objects.InitiateObjects(path)

	player.Position.X = playerX
	player.Position.Y = playerY
	scene.Background = background
	scene.ItemObjects = itemObjects
	scene.CollisionObjects = collisionObjects
	scene.CollisionObjects3d = collisionObjects3d
	scene.BloonObjects = bloonObjects

	CurrentScene = scene

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
}
