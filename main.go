package main

import (
	"BirdQuest/draw"
	"BirdQuest/global"
	"BirdQuest/menus"
	"BirdQuest/scene"
	"BirdQuest/update"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	player, camera := update.InitialLoader()
	if menus.AllMenus == nil {
		menus.PrepareMenus()
	}
	defer rl.CloseWindow()
	defer scene.UnloadAllBackgroundTextures()
	defer global.UnloadAllTextures()

	for !rl.WindowShouldClose() || rl.IsKeyPressed(rl.KeyEscape) {
		if scene.CurrentScene.Width == 0 {
			continue
		}

		if pause, p, c, exit := update.Menu(player, camera); pause || exit {
			if exit {
				break
			}

			if p != nil && c != nil {
				player = p
				camera = *c
			} else {
				update.Window(player, &camera)
				draw.Draw(camera, player)
				continue
			}
		}

		update.Update(&camera, player)

		update.Window(player, &camera)

		draw.Draw(camera, player)
	}
}
