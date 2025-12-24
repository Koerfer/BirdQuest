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
}

func Save(player *models.Player, camera rl.Camera2D) {
	state := &State{
		Player:          player,
		Camera:          camera,
		CurrentScene:    scene.CurrentScene,
		Scenes:          scene.AllScenes,
		GlobalVariables: global.VariableSet,
	}

	cwd, _ := os.Getwd()
	dumpFile, err := os.Create(cwd + "/save/save.bin")
	if err != nil {
		log.Fatalf("unable to create data.bin file: %v", err)
	}
	defer dumpFile.Close()

	enc := gob.NewEncoder(dumpFile)
	if err := enc.Encode(state); err != nil {
		log.Fatalf("failing to encode data: %v", err)
	}
}
