package global

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Variables struct {
	Fps      int32
	FpsScale float32
	Speed    float32

	EntityScale        float32
	EntitySize         float32
	PlayerMiddleOffset float32

	ScaleHeight float32
	ScaleWidth  float32

	DesiredHeight float32
	DesiredWidth  float32

	MapHeight float32
	MapWidth  float32

	VisibleMapHeight float32
	VisibleMapWidth  float32
}

var VariableSet *Variables

func SetDesiredWindowSize(width, height float32) {
	if VariableSet == nil {
		VariableSet = &Variables{}
	}
	rl.InitWindow(int32(width), int32(height), "BirdQuest")
	VariableSet.DesiredHeight = height
	VariableSet.DesiredWidth = width

	VariableSet.EntityScale = width / ScreenWidth
	VariableSet.ScaleHeight = height / ScreenHeight
	VariableSet.ScaleWidth = width / ScreenWidth

	if VariableSet.EntityScale < height/ScreenHeight {
		VariableSet.EntityScale = height / ScreenHeight
	}

	VariableSet.MapHeight = MapHeight * TileHeight * VariableSet.EntityScale
	VariableSet.MapWidth = MapWidth * TileWidth * VariableSet.EntityScale

	VariableSet.PlayerMiddleOffset = TileWidth / 2 * VariableSet.EntityScale
	VariableSet.EntitySize = TileWidth * VariableSet.EntityScale
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
