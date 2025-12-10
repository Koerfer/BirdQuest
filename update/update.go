package update

import (
	"BirdQuest/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var minZoom float32 = 1
var maxZoom float32 = 10

func Update(camera *rl.Camera2D, player *objects.Player, itemObjects, collisionObjects []*objects.Object, bloonObjects []*objects.Bloon) {
	updateZoom(camera, player)

	updatePlayer(camera, player, collisionObjects, bloonObjects)

	killItems(player, itemObjects, bloonObjects)
}
