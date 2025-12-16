package scene

import rl "github.com/gen2brain/raylib-go/raylib"

type Bloon struct {
	Object
	Lives                 int
	PoppingAnimationStage int
	AnimationStep         int
}

type Bloons struct {
	BloonObjects []*Bloon
	Texture      rl.Texture2D
}
