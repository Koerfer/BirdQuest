package update

import (
	"BirdQuest/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Update(camera *rl.Camera2D, player *objects.Player) {
	updateZoom(camera, player)

	updatePlayer(camera, player)

	killItems(player)
}
