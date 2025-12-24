package draw

import (
	"BirdQuest/global"
	"BirdQuest/scene"
	"BirdQuest/scene/models"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawDebugInfo(camera rl.Camera2D, player *models.Player) {
	//rl.DrawFPS(int32(camera.Target.X+5), int32(camera.Target.Y+5))

	for _, object := range scene.CurrentScene.ItemObjects.Objects {
		if object == nil {
			continue
		}
		rl.DrawRectanglePro(*object.Rectangle, rl.Vector2{X: 0, Y: 0}, 0, rl.Pink)
	}

	for _, object := range scene.CurrentScene.CollisionBoxes {
		if object == nil {
			continue
		}
		rl.DrawRectanglePro(*object, rl.Vector2{X: 0, Y: 0}, 0, rl.Red)
	}

	for _, bloon := range scene.CurrentScene.Bloons.BloonObjects {
		if bloon == nil {
			continue
		}
		rl.DrawRectanglePro(*bloon.Rectangle, rl.Vector2{X: 0, Y: 0}, 0, rl.Orange)
	}

	rl.DrawRectanglePro(*player.Rectangle, rl.Vector2{X: 0, Y: 0}, 0, rl.Green)

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
		player.Rectangle.X+global.VariableSet.PlayerMiddleOffset,
		player.Rectangle.Y+global.VariableSet.PlayerMiddleOffset,
	),
		int32(camera.Target.X+5),
		int32(camera.Target.Y+50/camera.Zoom),
		int32(40/camera.Zoom),
		rl.Black,
	)
	mousePositionRelative := rl.Vector2{
		X: mousePositionAbsolute.X/camera.Zoom + camera.Target.X - (player.Rectangle.X + global.VariableSet.PlayerMiddleOffset),
		Y: mousePositionAbsolute.Y/camera.Zoom + camera.Target.Y - (player.Rectangle.Y + global.VariableSet.PlayerMiddleOffset),
	}

	dashDirection := rl.Vector2Normalize(mousePositionRelative)
	dashDirection.X *= 100
	dashDirection.Y *= 100
	playerPositionVectorMiddle := rl.NewVector2(player.Rectangle.X+global.VariableSet.PlayerMiddleOffset, player.Rectangle.Y+global.VariableSet.PlayerMiddleOffset)
	rl.DrawLineV(playerPositionVectorMiddle, rl.Vector2Add(playerPositionVectorMiddle, dashDirection), rl.Black)

}
