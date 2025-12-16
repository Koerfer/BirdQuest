package update

import (
	"BirdQuest/scene"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Update(camera *rl.Camera2D, player *scene.Player) {
	updateZoom(camera, player)

	updatePlayer(camera, player)

	killItems(player)
}
