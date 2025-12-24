package models

import (
	"BirdQuest/global"
	rl "github.com/gen2brain/raylib-go/raylib"
	"time"
)

type Player struct {
	Object
	Animation     *Sprites
	IsMoving      bool
	AnimationStep int
	Rotation      float32
	DashCooldown  time.Duration
	DashLastUse   time.Time
	DashDirection rl.Vector2

	AttackCooldown time.Duration
	AttackLastUse  time.Time
	AttackOngoing  bool
}

func (player *Player) Draw() {
	var shadowOffset float32 = 20
	if !player.IsMoving && !player.AttackOngoing {
		shadowOffset = 2
	}

	rl.DrawTexturePro(
		global.VariableSet.PlayerTexture,
		rl.Rectangle{
			X:      player.BaseRectangle.X,
			Y:      player.BaseRectangle.Y + 96,
			Width:  player.BaseRectangle.Width,
			Height: player.BaseRectangle.Height,
		},
		rl.Rectangle{
			X:      player.Rectangle.X + global.VariableSet.PlayerMiddleOffset,
			Y:      player.Rectangle.Y + global.VariableSet.PlayerMiddleOffset + shadowOffset*global.VariableSet.EntityScale,
			Width:  player.Rectangle.Width,
			Height: player.Rectangle.Height,
		},
		rl.Vector2{X: global.VariableSet.PlayerMiddleOffset, Y: global.VariableSet.PlayerMiddleOffset},
		player.Rotation,
		rl.White,
	)

	rl.DrawTexturePro(
		global.VariableSet.PlayerTexture,
		*player.BaseRectangle,
		rl.Rectangle{
			X:      player.Rectangle.X + global.VariableSet.PlayerMiddleOffset,
			Y:      player.Rectangle.Y + global.VariableSet.PlayerMiddleOffset,
			Width:  player.Rectangle.Width,
			Height: player.Rectangle.Height,
		},
		rl.Vector2{X: global.VariableSet.PlayerMiddleOffset, Y: global.VariableSet.PlayerMiddleOffset},
		player.Rotation,
		rl.White,
	)
}
