package draw

import (
	"BirdQuest/scene"
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var debug = false

func Draw(camera rl.Camera2D, player *models.Player) {
	rl.BeginDrawing()

	rl.BeginMode2D(camera)

	drawBackground(camera)

	scene.CurrentScene.ItemObjects.Draw()
	scene.CurrentScene.Bloons.Draw()
	scene.CurrentScene.CollisionObjects.DrawFirstLayer()
	scene.CurrentScene.CollisionObjects.DrawDynamicLayer(player)
	scene.CurrentScene.CollisionObjects.DrawLastLayer()

	if debug {
		drawDebugInfo(camera, player)
	}

	rl.EndDrawing()
}
