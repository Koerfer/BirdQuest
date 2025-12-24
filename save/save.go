package save

import (
	"BirdQuest/scene"
	"BirdQuest/scene/models"
	"encoding/gob"
	"log"
	"os"
)

type State struct {
	Player *models.Player
	Scenes map[string]*models.Scene
}

func Save(player *models.Player) {
	state := &State{
		Player: player,
		Scenes: scene.AllScenes,
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
