package draw

import (
	"BirdQuest/global"
	"BirdQuest/objects"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawDebugInfo(camera rl.Camera2D, player *objects.Player, itemObjects, collisionObjects []*objects.Object, bloonObjects []*objects.Bloon) {
	//rl.DrawFPS(int32(camera.Target.X+5), int32(camera.Target.Y+5))

	for _, object := range itemObjects {
		if object == nil {
			continue
		}
		rl.DrawRectanglePro(object.HitBox, rl.Vector2{X: 0, Y: 0}, 0, rl.Pink)
	}

	for _, object := range collisionObjects {
		if object == nil {
			continue
		}
		rl.DrawRectanglePro(object.HitBox, rl.Vector2{X: 0, Y: 0}, 0, rl.Red)
	}

	for _, bloon := range bloonObjects {
		if bloon == nil {
			continue
		}
		rl.DrawRectanglePro(bloon.HitBox, rl.Vector2{X: 0, Y: 0}, 0, rl.Orange)
	}

	rl.DrawRectanglePro(player.HitBox, rl.Vector2{X: 0, Y: 0}, 0, rl.Green)

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
