package global

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Variables struct {
	Fps      int32
	FpsScale float32

	EntityScale            float32
	EntitySize             float32
	BasePlayerMiddleOffset float32
	PlayerMiddleOffset     float32

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
var Font rl.Font

func LoadAllTextures() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if time.Now().Month() == 12 && time.Now().Day() > 23 && time.Now().Day() < 31 {
		VariableSet.CollisionObjectsTexture = rl.LoadTexture(filepath.Join(cwd, "resources", "collision_sprites_winter.png"))
	} else {
		VariableSet.CollisionObjectsTexture = rl.LoadTexture(filepath.Join(cwd, "resources", "collision_sprites.png"))
	}
	VariableSet.BloonsTexture = rl.LoadTexture(filepath.Join(cwd, "resources", "bloons.png"))
	VariableSet.ItemsTexture = rl.LoadTexture(filepath.Join(cwd, "resources", "item_sprites.png"))
	VariableSet.PlayerTexture = rl.LoadTexture(filepath.Join(cwd, "resources", "chili.png"))

	Font = rl.LoadFontEx(filepath.Join(cwd, "resources", "fonts", "font.ttf"), 512, nil, 150)

}

func UnloadAllTextures() {
	rl.UnloadTexture(VariableSet.ItemsTexture)
	rl.UnloadTexture(VariableSet.CollisionObjectsTexture)
	rl.UnloadTexture(VariableSet.BloonsTexture)
	rl.UnloadTexture(VariableSet.PlayerTexture)

	rl.UnloadFont(Font)
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
}

func Zoom(zoom float32, camera *rl.Camera2D) {
	camera.Zoom = zoom
	VariableSet.VisibleMapHeight = VariableSet.DesiredHeight / zoom
	VariableSet.VisibleMapWidth = VariableSet.DesiredWidth / zoom
}
