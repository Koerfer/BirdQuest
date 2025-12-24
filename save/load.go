package save

import (
	"encoding/gob"
	"log"
	"os"
)

func Load() *State {
	cwd, _ := os.Getwd()
	_, err := os.Stat(cwd + "/save/save.bin")
	if os.IsNotExist(err) {
		return nil
	}

	binaryData, err := os.Open(cwd + "/save/save.bin")
	if err != nil {
		log.Fatalf("unable to open binary file: %v", err)
	}
	defer func(binaryData *os.File) {
		err := binaryData.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(binaryData)

	saveState := &State{}
	dec := gob.NewDecoder(binaryData)
	if err := dec.Decode(saveState); err != nil {
		log.Fatalf("failing to decode data: %v", err)
	}

	return saveState
}
