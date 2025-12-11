package update

import (
	"BirdQuest/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Update(camera *rl.Camera2D, player *objects.Player, itemObjects []*objects.Object, collisionObjects []*rl.Rectangle, bloonObjects []*objects.Bloon) {
	updateZoom(camera, player)

	updatePlayer(camera, player, collisionObjects, bloonObjects)

	killItems(player, itemObjects, bloonObjects)
}
