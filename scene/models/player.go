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

	Talking      bool
	DialogStep   int
	DialogNPC    *NPC
	CurrentQuest *Quest
}

func (player *Player) Draw() {
	var shadowOffset float32 = 20
	if !player.IsMoving && !player.AttackOngoing {
		shadowOffset = 2
	}

	rl.DrawTexturePro(
		global.VariableSet.Textures32x32,
		rl.Rectangle{
			X:      player.BaseRectangle.X,
			Y:      player.BaseRectangle.Y + 32,
			Width:  player.BaseRectangle.Width,
			Height: player.BaseRectangle.Height,
		},
		rl.Rectangle{
			X:      player.BasePositionRectangle.X*global.VariableSet.EntityScale + global.VariableSet.PlayerMiddleOffset,
			Y:      player.BasePositionRectangle.Y*global.VariableSet.EntityScale + global.VariableSet.PlayerMiddleOffset + shadowOffset*global.VariableSet.EntityScale,
			Width:  player.BasePositionRectangle.Width * global.VariableSet.EntityScale,
			Height: player.BasePositionRectangle.Height * global.VariableSet.EntityScale,
		},
		rl.Vector2{X: global.VariableSet.PlayerMiddleOffset, Y: global.VariableSet.PlayerMiddleOffset},
		player.Rotation,
		rl.White,
	)

	rl.DrawTexturePro(
		global.VariableSet.Textures32x32,
		*player.BaseRectangle,
		rl.Rectangle{
			X:      player.BasePositionRectangle.X*global.VariableSet.EntityScale + global.VariableSet.PlayerMiddleOffset,
			Y:      player.BasePositionRectangle.Y*global.VariableSet.EntityScale + global.VariableSet.PlayerMiddleOffset,
			Width:  player.BasePositionRectangle.Width * global.VariableSet.EntityScale,
			Height: player.BasePositionRectangle.Height * global.VariableSet.EntityScale,
		},
		rl.Vector2{X: global.VariableSet.PlayerMiddleOffset, Y: global.VariableSet.PlayerMiddleOffset},
		player.Rotation,
		rl.White,
	)
}
