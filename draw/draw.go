package draw

import (
	"BirdQuest/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var debug = false

func Draw(camera rl.Camera2D, player *objects.Player) {
	rl.BeginDrawing()

	rl.BeginMode2D(camera)

	drawBackground(camera)

	drawObjects()

	drawBloons()

	drawAfterPlayer := drawCollisionObjects(player)

	drawPlayer(player)

	drawCollisionObjectsAfterPlayer(drawAfterPlayer)

	if debug {
		drawDebugInfo(camera, player)
	}

	rl.EndDrawing()
}
