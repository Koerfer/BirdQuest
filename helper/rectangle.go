package helper

import rl "github.com/gen2brain/raylib-go/raylib"

func MultiplyRectangle(rectangle *rl.Rectangle, factor float32) rl.Rectangle {
	return rl.Rectangle{
		X:      rectangle.X * factor,
		Y:      rectangle.Y * factor,
		Width:  rectangle.Width * factor,
		Height: rectangle.Height * factor,
	}
}
