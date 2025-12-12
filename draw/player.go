package draw

import (
	"BirdQuest/global"
	"BirdQuest/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawPlayer(player *objects.Player) {
	var shadowOffset float32 = 20
	if !player.IsMoving && !player.AttackOngoing {
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
			X:      player.Position.X + global.VariableSet.PlayerMiddleOffset,
			Y:      player.Position.Y + global.VariableSet.PlayerMiddleOffset + shadowOffset*global.VariableSet.EntityScale,
			Width:  player.Rectangle.Width * global.VariableSet.EntityScale,
			Height: player.Rectangle.Height * global.VariableSet.EntityScale,
		},
		rl.Vector2{X: global.VariableSet.PlayerMiddleOffset, Y: global.VariableSet.PlayerMiddleOffset},
		player.Rotation,
		rl.White,
	)

	rl.DrawTexturePro(
		player.Texture,
		player.Rectangle,
		rl.Rectangle{
			X:      player.Position.X + global.VariableSet.PlayerMiddleOffset,
			Y:      player.Position.Y + global.VariableSet.PlayerMiddleOffset,
			Width:  player.Rectangle.Width * global.VariableSet.EntityScale,
			Height: player.Rectangle.Height * global.VariableSet.EntityScale,
		},
		rl.Vector2{X: global.VariableSet.PlayerMiddleOffset, Y: global.VariableSet.PlayerMiddleOffset},
		player.Rotation,
		rl.White,
	)
}
