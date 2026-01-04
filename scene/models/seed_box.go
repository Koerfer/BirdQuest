package models

import (
	"BirdQuest/global"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type SeedBox struct {
	OpeningStage float32

	Object
}

func (box *SeedBox) Draw() {
	if box == nil {
		return
	}

	rl.DrawTexturePro(
		global.VariableSet.Textures32x32,
		*box.BaseRectangle,
		rl.Rectangle{
			X:      box.BasePositionRectangle.X * global.VariableSet.EntityScale,
			Y:      box.BasePositionRectangle.Y * global.VariableSet.EntityScale,
			Width:  box.BasePositionRectangle.Width * global.VariableSet.EntityScale,
			Height: box.BasePositionRectangle.Height * global.VariableSet.EntityScale,
		},
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)
}
