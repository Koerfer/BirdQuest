package global

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
	"os"
	"path/filepath"
)

type Variables struct {
	Fps      int32
	FpsScale float32
	Speed    float32

	EntityScale            float32
	EntitySize             float32
	BasePlayerMiddleOffset float32
	PlayerMiddleOffset     float32

	ScaleHeight float32
	ScaleWidth  float32

	DesiredHeight float32
	DesiredWidth  float32

	MapHeight float32
	MapWidth  float32

	VisibleMapHeight float32
	VisibleMapWidth  float32

	ItemsTexture            rl.Texture2D
	BloonsTexture           rl.Texture2D
	CollisionObjectsTexture rl.Texture2D
	PlayerTexture           rl.Texture2D
}

var VariableSet *Variables

func LoadAllTextures() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	VariableSet.CollisionObjectsTexture = rl.LoadTexture(filepath.Join(cwd, "sprites", "collision_sprites.png"))
	VariableSet.BloonsTexture = rl.LoadTexture(filepath.Join(cwd, "sprites", "bloons.png"))
	VariableSet.ItemsTexture = rl.LoadTexture(filepath.Join(cwd, "sprites", "item_sprites.png"))
	VariableSet.PlayerTexture = rl.LoadTexture(filepath.Join(cwd, "sprites", "chili.png"))
}

func UnloadAllTextures() {
	rl.UnloadTexture(VariableSet.ItemsTexture)
	rl.UnloadTexture(VariableSet.CollisionObjectsTexture)
	rl.UnloadTexture(VariableSet.BloonsTexture)
	rl.UnloadTexture(VariableSet.PlayerTexture)
}

func SetDesiredWindowSize(width, height float32) {
	if VariableSet == nil {
		VariableSet = &Variables{}
	}
	rl.InitWindow(int32(width), int32(height), "BirdQuest")
	rl.SetWindowState(rl.FlagWindowResizable)
	VariableSet.DesiredHeight = height
	VariableSet.DesiredWidth = width
}

func SetFPS(fps int32) {
	if VariableSet == nil {
		VariableSet = &Variables{}
	}

	rl.SetTargetFPS(fps)

	VariableSet.Fps = fps
	VariableSet.FpsScale = 60 / float32(fps)
	VariableSet.Speed = VariableSet.FpsScale * VariableSet.EntityScale
}

func Zoom(zoom float32, camera *rl.Camera2D) {
	camera.Zoom = zoom
	VariableSet.VisibleMapHeight = VariableSet.DesiredHeight / zoom
	VariableSet.VisibleMapWidth = VariableSet.DesiredWidth / zoom
}
