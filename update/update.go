package update

import (
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Update(camera *rl.Camera2D, player *models.Player) {
	updateZoom(camera, player)

	updatePlayer(camera, player)

	killItems(player)
}
