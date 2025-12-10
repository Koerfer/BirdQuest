package draw

import (
	"BirdQuest/global"
	"BirdQuest/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawObjects(objects []*objects.Object) {
	for _, object := range objects {
		drawObject(object)
	}
}

func drawObject(object *objects.Object) {
	if object == nil {
		return
	}

	rl.DrawTexturePro(
		object.Texture,
		object.Rectangle,
		rl.Rectangle{
			X:      object.Position.X,
			Y:      object.Position.Y,
			Width:  object.Rectangle.Width * global.Scale,
			Height: object.Rectangle.Height * global.Scale,
		},
		rl.Vector2{
			X: 0,
			Y: 0,
		},
		0,
		rl.White,
	)
}
