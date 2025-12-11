package draw

import (
	"BirdQuest/objects"
)

func drawCollisionObjects(collisionObjects []*objects.Object, player *objects.Player) []*objects.Object {
	var drawAfterPlayer []*objects.Object

	for _, collisionObject := range collisionObjects {
		if collisionObject.AlwaysRenderLast {
			if drawAfterPlayer == nil {
				drawAfterPlayer = make([]*objects.Object, 0)
			}

			drawAfterPlayer = append(drawAfterPlayer, collisionObject)
			continue
		}

		if collisionObject.Position.Y >= player.Position.Y {
			if drawAfterPlayer == nil {
				drawAfterPlayer = make([]*objects.Object, 0)
			}

			drawAfterPlayer = append(drawAfterPlayer, collisionObject)
			continue
		}

		drawObject(collisionObject)
	}

	return drawAfterPlayer
}

func drawCollisionObjectsAfterPlayer(collisionObjects []*objects.Object) {
	if collisionObjects == nil {
		return
	}

	for _, object := range collisionObjects {
		drawObject(object)
	}
}
