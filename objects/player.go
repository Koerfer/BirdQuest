package objects

import (
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
