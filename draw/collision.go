package draw

import (
	"BirdQuest/scene"
)

func drawCollisionObjects(player *scene.Player) []*scene.Object {
	var drawAfterPlayer []*scene.Object

	for _, collisionObject := range scene.CurrentScene.CollisionObjects3d {
		if collisionObject.AlwaysRenderFirst {
			drawObject(collisionObject)
			continue
		}

		if collisionObject.AlwaysRenderLast {
			if drawAfterPlayer == nil {
				drawAfterPlayer = make([]*scene.Object, 0)
			}

			drawAfterPlayer = append(drawAfterPlayer, collisionObject)
			continue
		}

		if collisionObject.Position.Y >= player.Position.Y {
			if drawAfterPlayer == nil {
				drawAfterPlayer = make([]*scene.Object, 0)
			}

			drawAfterPlayer = append(drawAfterPlayer, collisionObject)
			continue
		}

		drawObject(collisionObject)
	}

	return drawAfterPlayer
}

func drawCollisionObjectsAfterPlayer(collisionObjects []*scene.Object) {
	if collisionObjects == nil {
		return
	}

	for _, object := range collisionObjects {
		drawObject(object)
	}
}
