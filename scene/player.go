package scene

import (
	"BirdQuest/global"
	rl "github.com/gen2brain/raylib-go/raylib"
	"path/filepath"
	"time"
)

type Player struct {
	Object
	Animation     *sprites
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

func preparePlayer(path string) *Player {
	chiliAnimationsRaw := rl.LoadTexture(filepath.Join(path, "chili.png"))

	chiliAnimations := &sprites{
		Texture:      chiliAnimationsRaw,
		TileWidth:    global.TileWidth,
		TileHeight:   global.TileHeight,
		WidthInTiles: int(chiliAnimationsRaw.Width) / global.TileWidth,
	}

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
			Position:  rl.Vector2{},
			Texture:   chiliAnimations.Texture,
			Rectangle: chiliAnimations.GetRectangleAreaInTexture(7),
			HitBox: rl.Rectangle{
				X:      0,
				Y:      0,
				Width:  global.VariableSet.EntitySize,
				Height: global.VariableSet.EntitySize,
			},
		}}
}
