package main

import (
	"BirdQuest/draw"
	"BirdQuest/global"
	objectsPkg "BirdQuest/objects"
	"BirdQuest/update"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
	"os"
	"path/filepath"
)

func main() {
	rl.InitWindow(global.ScreenWidth*global.Scale, global.ScreenHeight*global.Scale, "BirdQuest")
	defer rl.CloseWindow()

	rl.SetTargetFPS(144)

	camera := rl.Camera2D{}
	camera.Target = rl.Vector2{
		X: 0,
		Y: 0,
	}
	camera.Zoom = 3

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	backgroundRaw := rl.LoadTexture(filepath.Join(cwd, "sprites/background.png"))
	defer rl.UnloadTexture(backgroundRaw)
	collisionSpritesRaw := rl.LoadTexture(filepath.Join(cwd, "sprites/collision_sprites.png"))
	defer rl.UnloadTexture(collisionSpritesRaw)
	itemSpritesRaw := rl.LoadTexture(filepath.Join(cwd, "sprites/item_sprites.png"))
	defer rl.UnloadTexture(itemSpritesRaw)
	chiliAnimationsRaw := rl.LoadTexture(filepath.Join(cwd, "sprites/chili.png"))
	defer rl.UnloadTexture(chiliAnimationsRaw)
	bloonsSpritesRaw := rl.LoadTexture(filepath.Join(cwd, "sprites/bloons.png"))
	defer rl.UnloadTexture(bloonsSpritesRaw)

	itemObjects, collisionObjects, bloonObjects, player := objectsPkg.InitiateObjects(cwd, collisionSpritesRaw, itemSpritesRaw, bloonsSpritesRaw, chiliAnimationsRaw)

	for !rl.WindowShouldClose() { // Detect window close button or ESC key
		update.Update(&camera, player, itemObjects, collisionObjects, bloonObjects)

		draw.Draw(camera, itemObjects, collisionObjects, bloonObjects, player, backgroundRaw)
	}
}
