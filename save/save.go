package save

import (
	"BirdQuest/global"
	"BirdQuest/scene"
	"BirdQuest/scene/models"
	"encoding/gob"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
	"os"
)

type State struct {
	Player          *models.Player
	Camera          rl.Camera2D
	CurrentScene    *models.Scene
	Scenes          map[string]*models.Scene
	GlobalVariables *global.Variables

	IsFullScreen   bool
	IsMaximised    bool
	WindowWidth    float32
	WindowHeight   float32
	WindowPosition rl.Vector2
}

func Save(player *models.Player, camera rl.Camera2D) {
	state := &State{
		Player:          player,
		Camera:          camera,
		CurrentScene:    scene.CurrentScene,
		Scenes:          scene.AllScenes,
		GlobalVariables: global.VariableSet,

		IsFullScreen:   rl.IsWindowFullscreen(),
		IsMaximised:    rl.IsWindowMaximized(),
		WindowWidth:    float32(rl.GetScreenWidth()),
		WindowHeight:   float32(rl.GetScreenHeight()),
		WindowPosition: rl.GetWindowPosition(),
	}

	cwd, _ := os.Getwd()
	dumpFile, err := os.Create(cwd + "/save/save.bin")
	if err != nil {
		log.Fatalf("unable to create data.bin file: %v", err)
	}
	defer func(dumpFile *os.File) {
		err := dumpFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dumpFile)

	enc := gob.NewEncoder(dumpFile)
	if err := enc.Encode(state); err != nil {
		log.Fatalf("failing to encode data: %v", err)
	}
}
