package draw

import (
	"BirdQuest/menus"
	"BirdQuest/scene"
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Draw(camera rl.Camera2D, player *models.Player) {
	rl.BeginDrawing()

	rl.BeginMode2D(camera)

	drawBackground(camera)

	scene.CurrentScene.ItemObjects.Draw()
	scene.CurrentScene.Bloons.Draw()
	scene.CurrentScene.CollisionObjects.DrawFirstLayer()
	scene.CurrentScene.CollisionObjects.DrawDynamicLayer(player)
	scene.CurrentScene.CollisionObjects.DrawLastLayer()

	for _, npc := range scene.CurrentScene.NPCs {
		npc.Draw()
	}

	//drawDebugInfo(camera, player)

	if player.Talking {
		player.CurrentQuest.Steps[player.CurrentQuest.CurrentStep].Dialogs[player.DialogStep].Draw(camera, player.DialogNPC, player)
	}

	if menus.ActiveMenu != nil {
		menus.ActiveMenu.Draw(camera)
	}

	rl.EndDrawing()
}
