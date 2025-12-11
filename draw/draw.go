package draw

import (
	"BirdQuest/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var debug = false

func Draw(
	camera rl.Camera2D,
	itemObjects []*objects.Object,
	collisionObjects []*rl.Rectangle,
	collisionObjects3d []*objects.Object,
	bloonObjects []*objects.Bloon,
	player *objects.Player,
	backgroundRaw rl.Texture2D) {
	rl.BeginDrawing()

	rl.BeginMode2D(camera)

	drawBackground(camera, backgroundRaw)

	drawObjects(itemObjects)

	drawBloons(bloonObjects)

	drawAfterPlayer := drawCollisionObjects(collisionObjects3d, player)

	drawPlayer(player)

	drawCollisionObjectsAfterPlayer(drawAfterPlayer)

	if debug {
		drawDebugInfo(camera, player, itemObjects, collisionObjects, bloonObjects)
	}

	rl.EndDrawing()
}
