package main

import (
	"BirdQuest/draw"
	"BirdQuest/global"
	"BirdQuest/scene"
	"BirdQuest/update"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	player, camera := update.InitialLoader()
	defer rl.CloseWindow()
	defer scene.UnloadAllBackgroundTextures()
	defer global.UnloadAllTextures()

	for !rl.WindowShouldClose() || rl.IsKeyPressed(rl.KeyEscape) {
		if scene.CurrentScene.Width == 0 {
			continue
		}

		if p, c := update.SaveHandler(player, &camera); p != nil && c != nil {
			player = p
			camera = *c
		}

		update.Update(&camera, player)

		update.Window(player, &camera)

		draw.Draw(camera, player)
	}
}
