package models

import rl "github.com/gen2brain/raylib-go/raylib"

type Object struct {
	BasePosition  *rl.Vector2
	BaseRectangle *rl.Rectangle

	Rectangle *rl.Rectangle
}
