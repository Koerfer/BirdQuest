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
}
