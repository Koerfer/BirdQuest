package models

import rl "github.com/gen2brain/raylib-go/raylib"

type Door struct {
	BaseRectangle *rl.Rectangle

	GoesToScene string
	GoesToX     float32
	GoesToY     float32
}
