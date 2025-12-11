package objects

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

func preparePlayer(chiliAnimations *Sprites) *Player {
	return &Player{
		IsMoving:       false,
		AnimationStep:  0,
		Rotation:       0,
		Animation:      chiliAnimations,
		DashLastUse:    time.Time{},
		DashCooldown:   time.Millisecond * 1200,
		AttackLastUse:  time.Time{},
		AttackCooldown: time.Millisecond * 500,
		Object: Object{
			Position:  rl.Vector2{X: global.PlayerStartX * global.Scale, Y: global.PlayerStartY * global.Scale},
			Texture:   chiliAnimations.Texture,
			Rectangle: chiliAnimations.GetSrc(7),
			HitBox: rl.Rectangle{
				X:      0 * global.Scale,
				Y:      0 * global.Scale,
				Width:  float32(global.TileWidth) * global.Scale,
				Height: float32(global.TileWidth) * global.Scale,
			},
		}}
}
