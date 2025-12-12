package draw

import (
	"BirdQuest/global"
	"BirdQuest/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawBloons(bloons []*objects.Bloon) {
	for _, bloon := range bloons {
		if bloon == nil {
			continue
		}

		// If bloon is popped
		if bloon.PoppingAnimationStage > 0 {
			bloon.Rectangle = rl.Rectangle{
				X:      global.TileWidth * (float32(bloon.PoppingAnimationStage + 2)),
				Y:      0,
				Width:  bloon.Rectangle.Width,
				Height: bloon.Rectangle.Height,
			}
		}

		drawObject(&bloon.Object)
	}
}
