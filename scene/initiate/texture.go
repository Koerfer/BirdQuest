package initiate

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"path/filepath"
)

func prepareTexture(path, pngName string) rl.Texture2D {
	return rl.LoadTexture(filepath.Join(path, pngName))
}
