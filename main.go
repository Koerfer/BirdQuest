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

	player := scene.SetScene("main", 250, 250, nil)
	movement.InitialiseCamera(player, &camera)

	for !rl.WindowShouldClose() {
		if scene.CurrentScene.Width == 0 {
			continue
		}

		update.SaveHandler(player, camera)

		update.Update(&camera, player)

		update.Window(player, &camera)

		draw.Draw(camera, player)
	}

	scene.UnloadAllTextures()
}
