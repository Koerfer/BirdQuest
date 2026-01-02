package models

import (
	"BirdQuest/global"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Items struct {
	Objects []*Object
}

func (items *Items) Draw() {
	for _, item := range items.Objects {
		if item == nil {
			continue
		}

		rl.DrawTexturePro(
			global.VariableSet.Textures32x32,
			*item.BaseRectangle,
			rl.Rectangle{
				X:      item.BasePositionRectangle.X * global.VariableSet.EntityScale,
				Y:      item.BasePositionRectangle.Y * global.VariableSet.EntityScale,
				Width:  item.BasePositionRectangle.Width * global.VariableSet.EntityScale,
				Height: item.BasePositionRectangle.Height * global.VariableSet.EntityScale,
			},
			rl.Vector2{
				X: 0,
				Y: 0,
			},
			0,
			rl.White,
		)
	}
}
