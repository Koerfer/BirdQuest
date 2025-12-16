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

	player := scene.SetScene("", 250, 250)
	movement.InitialiseCamera(player, &camera)

	for !rl.WindowShouldClose() { // Detect window close button or ESC key
		update.Update(&camera, player)

		draw.Draw(camera, player)
	}

	scene.UnloadAllTextures()
}
