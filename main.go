package main

import (
	"BirdQuest/draw"
	"BirdQuest/global"
	"BirdQuest/movement"
	"BirdQuest/scene"
	"BirdQuest/update"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	global.SetDesiredWindowSize(1920, 1080)
	defer rl.CloseWindow()

	global.SetFPS(120)

	camera := rl.Camera2D{}
	camera.Target = rl.Vector2{}
	global.Zoom(1, &camera)

	player := scene.SetScene("house", 250, 250)
	movement.InitialiseCamera(player, &camera)

	for !rl.WindowShouldClose() {
		if scene.CurrentScene.Width == 0 {
			continue
		}

		update.Update(&camera, player)

		if rl.IsKeyPressed(rl.KeyOne) {
			switch scene.CurrentScene.Name {
			case "main":
				scene.UnloadAllTextures()
				player = scene.SetScene("house", 250, 250)
			case "house":
				scene.UnloadAllTextures()
				player = scene.SetScene("main", 250, 250)
			}
			movement.InitialiseCamera(player, &camera)
		}

		draw.Draw(camera, player)
	}

	scene.UnloadAllTextures()
}
