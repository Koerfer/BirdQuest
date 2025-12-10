package draw

import (
	"BirdQuest/global"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawBackground(camera rl.Camera2D, backgroundRaw rl.Texture2D) {
	rl.ClearBackground(rl.Black)

	drawSizeWidth := global.ScreenWidth/camera.Zoom + 20
	drawSizeHeight := global.ScreenHeight/camera.Zoom + 20
	if drawSizeWidth > global.ScreenWidth {
		drawSizeWidth = global.ScreenWidth
	}
	if drawSizeHeight > global.ScreenHeight {
		drawSizeHeight = global.ScreenHeight
	}
	rl.DrawTexturePro(
		backgroundRaw,
		rl.Rectangle{
			X:      camera.Target.X / global.Scale,
			Y:      camera.Target.Y / global.Scale,
			Width:  camera.Target.X/global.Scale + drawSizeWidth,
			Height: camera.Target.Y/global.Scale + drawSizeHeight,
		},
		rl.Rectangle{
			X:      camera.Target.X,
			Y:      camera.Target.Y,
			Width:  camera.Target.X + drawSizeWidth*global.Scale,
			Height: camera.Target.Y + drawSizeHeight*global.Scale,
		},
		rl.Vector2{
			X: 0,
			Y: 0,
		},
		0,
		rl.White,
	)
}
