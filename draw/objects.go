package draw

import (
	"BirdQuest/global"
	"BirdQuest/objects"
	"BirdQuest/scene"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawObjects() {
	for _, object := range scene.CurrentScene.ItemObjects {
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
			Width:  object.Rectangle.Width * global.VariableSet.EntityScale,
			Height: object.Rectangle.Height * global.VariableSet.EntityScale,
		},
		rl.Vector2{
			X: 0,
			Y: 0,
		},
		0,
		rl.White,
	)
}
