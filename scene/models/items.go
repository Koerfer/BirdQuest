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
			global.VariableSet.ItemsTexture,
			*item.BaseRectangle,
			rl.Rectangle{
				X:      item.Rectangle.X,
				Y:      item.Rectangle.Y,
				Width:  item.Rectangle.Width,
				Height: item.Rectangle.Height,
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
