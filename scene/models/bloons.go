package models

import (
	"BirdQuest/global"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Bloon struct {
	Object
	Lives                 int
	PoppingAnimationStage int
	AnimationStep         int
}

type Bloons struct {
	BloonObjects []*Bloon
}

func (bloons *Bloons) Draw() {

	for _, bloon := range bloons.BloonObjects {
		if bloon == nil {
			continue
		}

		// If bloon is popped
		if bloon.PoppingAnimationStage > 0 {
			bloon.BaseRectangle = &rl.Rectangle{
				X:      global.TileWidth * (float32(bloon.PoppingAnimationStage + 2)),
				Y:      0,
				Width:  bloon.BaseRectangle.Width,
				Height: bloon.BaseRectangle.Height,
			}
		}

		rl.DrawTexturePro(
			global.VariableSet.BloonsTexture,
			*bloon.BaseRectangle,
			rl.Rectangle{
				X:      bloon.Rectangle.X,
				Y:      bloon.Rectangle.Y,
				Width:  bloon.Rectangle.Width,
				Height: bloon.Rectangle.Height,
			},
			rl.Vector2{
				X: 0,
				Y: 0,
			},
			0,
			rl.White,
		)
	}

}
