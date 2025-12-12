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
			X:      camera.Target.X / global.VariableSet.EntityScale,
			Y:      camera.Target.Y / global.VariableSet.EntityScale,
			Width:  camera.Target.X/global.VariableSet.EntityScale + drawSizeWidth,
			Height: camera.Target.Y/global.VariableSet.EntityScale + drawSizeHeight,
		},
		rl.Rectangle{
			X:      camera.Target.X,
			Y:      camera.Target.Y,
			Width:  camera.Target.X + drawSizeWidth*global.VariableSet.EntityScale,
			Height: camera.Target.Y + drawSizeHeight*global.VariableSet.EntityScale,
		},
		rl.Vector2{
			X: 0,
			Y: 0,
		},
		0,
		rl.White,
	)
}
