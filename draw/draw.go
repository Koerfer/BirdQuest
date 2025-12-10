package draw

import (
	"BirdQuest/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var debug = true

func Draw(camera rl.Camera2D, itemObjects, collisionObjects []*objects.Object, bloonObjects []*objects.Bloon, player *objects.Player, backgroundRaw rl.Texture2D) {
	rl.BeginDrawing()

	rl.BeginMode2D(camera)

	drawBackground(camera, backgroundRaw)

	if debug {
		drawDebugInfo(camera, player, itemObjects, collisionObjects, bloonObjects)
	}

	drawObjects(itemObjects)

	drawBloons(bloonObjects)

	drawPlayer(player)

	drawObjects(collisionObjects)

	rl.EndDrawing()
}
