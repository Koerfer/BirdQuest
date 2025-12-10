package draw

import (
	"BirdQuest/global"
	"BirdQuest/objects"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"slices"
)

var debug = false

func Draw(camera rl.Camera2D, itemObjects, collisionObjects []*objects.Object, bloonObjects []*objects.Bloon, player *objects.Player, backgroundRaw rl.Texture2D) {
	rl.BeginDrawing()

	rl.BeginMode2D(camera)

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

	if debug {
		rl.DrawRectanglePro(player.HitBox, rl.Vector2{X: 0, Y: 0}, 0, rl.Pink)
	}

	for _, object := range itemObjects {
		if object == nil {
			continue
		}
		if debug {
			rl.DrawRectanglePro(object.HitBox, rl.Vector2{X: 0, Y: 0}, 0, rl.Pink)
		}

		rl.DrawTexturePro(
			object.Texture,
			object.Rectangle,
			rl.Rectangle{
				X:      object.Position.X,
				Y:      object.Position.Y,
				Width:  object.Rectangle.Width * global.Scale,
				Height: object.Rectangle.Height * global.Scale,
			},
			rl.Vector2{
				X: 0,
				Y: 0,
			},
			0,
			rl.White,
		)

	}

	for i, object := range bloonObjects {
		if object == nil {
			continue
		}

		if debug {
			rl.DrawRectanglePro(object.HitBox, rl.Vector2{X: 0, Y: 0}, 0, rl.Pink)
		}

		if object.PoppingAnimationStage > 0 {
			object.AnimationStep++
			if object.AnimationStep == int(70*60/rl.GetFPS()) {
				object.AnimationStep = 0
				object.PoppingAnimationStage++
			}

			if object.PoppingAnimationStage == 5 {
				bloonObjects = slices.Delete(bloonObjects, i, i+1)
				continue
			}

			rl.DrawTexturePro(
				object.Texture,
				rl.Rectangle{
					X:      32 * (float32(object.PoppingAnimationStage + 2)),
					Y:      0,
					Width:  object.Rectangle.Width,
					Height: object.Rectangle.Height,
				},
				rl.Rectangle{
					X:      object.Position.X,
					Y:      object.Position.Y,
					Width:  object.Rectangle.Width * global.Scale,
					Height: object.Rectangle.Height * global.Scale,
				},
				rl.Vector2{
					X: 0,
					Y: 0,
				},
				0,
				rl.White,
			)
			continue
		}

		rl.DrawTexturePro(
			object.Texture,
			object.Rectangle,
			rl.Rectangle{
				X:      object.Position.X,
				Y:      object.Position.Y,
				Width:  object.Rectangle.Width * global.Scale,
				Height: object.Rectangle.Height * global.Scale,
			},
			rl.Vector2{
				X: 0,
				Y: 0,
			},
			0,
			rl.White,
		)

	}

	var shadowOffset float32 = 20
	if !player.IsMoving {
		shadowOffset = 2
	}
	rl.DrawTexturePro(
		player.Texture,
		rl.Rectangle{
			X:      player.Rectangle.X,
			Y:      player.Rectangle.Y + 96,
			Width:  player.Rectangle.Width,
			Height: player.Rectangle.Height,
		},
		rl.Rectangle{
			X:      player.Position.X + 16*global.Scale,
			Y:      player.Position.Y + 16*global.Scale + shadowOffset*global.Scale,
			Width:  player.Rectangle.Width * global.Scale,
			Height: player.Rectangle.Height * global.Scale,
		},
		rl.Vector2{X: 16 * global.Scale, Y: 16 * global.Scale},
		player.Rotation,
		rl.White,
	)

	rl.DrawTexturePro(
		player.Texture,
		player.Rectangle,
		rl.Rectangle{
			X:      player.Position.X + 16*global.Scale,
			Y:      player.Position.Y + 16*global.Scale,
			Width:  player.Rectangle.Width * global.Scale,
			Height: player.Rectangle.Height * global.Scale,
		},
		rl.Vector2{X: 16 * global.Scale, Y: 16 * global.Scale},
		player.Rotation,
		rl.White,
	)

	for _, object := range collisionObjects {
		if object == nil {
			continue
		}
		if debug {
			rl.DrawRectanglePro(object.HitBox, rl.Vector2{X: 0, Y: 0}, 0, rl.Pink)
		}

		rl.DrawTexturePro(
			object.Texture,
			object.Rectangle,
			rl.Rectangle{
				X:      object.Position.X,
				Y:      object.Position.Y,
				Width:  object.Rectangle.Width * global.Scale,
				Height: object.Rectangle.Height * global.Scale,
			},
			rl.Vector2{
				X: 0,
				Y: 0,
			},
			0,
			rl.White,
		)
	}

	if debug {
		//rl.DrawFPS(int32(camera.Target.X+5), int32(camera.Target.Y+5))

		mousePositionAbsolute := rl.GetMousePosition()
		rl.DrawText(fmt.Sprintf("%f, %f",
			mousePositionAbsolute.X/2+camera.Target.X,
			mousePositionAbsolute.Y/2+camera.Target.Y,
		),
			int32(camera.Target.X+5),
			int32(camera.Target.Y+5/camera.Zoom),
			int32(40/camera.Zoom),
			rl.Black,
		)
		rl.DrawText(fmt.Sprintf("%f, %f",
			player.Position.X+16*global.Scale,
			player.Position.Y+16*global.Scale,
		),
			int32(camera.Target.X+5),
			int32(camera.Target.Y+50/camera.Zoom),
			int32(40/camera.Zoom),
			rl.Black,
		)
		mousePositionRelative := rl.Vector2{
			X: mousePositionAbsolute.X/camera.Zoom + camera.Target.X - (player.Position.X + 16*global.Scale),
			Y: mousePositionAbsolute.Y/camera.Zoom + camera.Target.Y - (player.Position.Y + 16*global.Scale),
		}

		dashDirection := rl.Vector2Normalize(mousePositionRelative)
		dashDirection.X *= 100
		dashDirection.Y *= 100
		playerPositionVectorMiddle := rl.NewVector2(player.Position.X+16*global.Scale, player.Position.Y+16*global.Scale)
		rl.DrawLineV(playerPositionVectorMiddle, rl.Vector2Add(playerPositionVectorMiddle, dashDirection), rl.Black)
	}

	rl.EndDrawing()
}
