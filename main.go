package main

import (
	"BirdQuest/draw"
	"BirdQuest/global"
	"BirdQuest/movement"
	objectsPkg "BirdQuest/objects"
	"BirdQuest/update"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
	"os"
	"path/filepath"
)

func main() {
	global.SetDesiredSize(1920, 1080)
	defer rl.CloseWindow()

	global.SetFPS(120)

	camera := rl.Camera2D{}
	camera.Target = rl.Vector2{}
	global.Zoom(3, &camera)

	fmt.Println(global.VariableSet)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	backgroundRaw := rl.LoadTexture(filepath.Join(cwd, "sprites/background.png"))
	defer rl.UnloadTexture(backgroundRaw)
	itemSpritesRaw := rl.LoadTexture(filepath.Join(cwd, "sprites/item_sprites.png"))
	defer rl.UnloadTexture(itemSpritesRaw)
	collisionSpritesRaw := rl.LoadTexture(filepath.Join(cwd, "sprites/collision_sprites.png"))
	defer rl.UnloadTexture(collisionSpritesRaw)
	chiliAnimationsRaw := rl.LoadTexture(filepath.Join(cwd, "sprites/chili.png"))
	defer rl.UnloadTexture(chiliAnimationsRaw)
	bloonsSpritesRaw := rl.LoadTexture(filepath.Join(cwd, "sprites/bloons.png"))
	defer rl.UnloadTexture(bloonsSpritesRaw)

	itemObjects, collisionObjects, collisionObjects3d, bloonObjects, player := objectsPkg.InitiateObjects(cwd, itemSpritesRaw, bloonsSpritesRaw, chiliAnimationsRaw, collisionSpritesRaw)
	movement.InitialiseCamera(player, &camera)

	for !rl.WindowShouldClose() { // Detect window close button or ESC key
		update.Update(&camera, player, itemObjects, collisionObjects, bloonObjects)

		draw.Draw(camera, itemObjects, collisionObjects, collisionObjects3d, bloonObjects, player, backgroundRaw)
	}
}
