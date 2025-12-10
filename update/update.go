package update

import (
	"BirdQuest/attack"
	"BirdQuest/global"
	"BirdQuest/movement"
	"BirdQuest/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
	"slices"
	"time"
)

var minZoom float32 = 1
var maxZoom float32 = 10

func Update(camera *rl.Camera2D, player *objects.Player, itemObjects, collisionObjects []*objects.Object, bloonObjects []*objects.Bloon) {
	fps := rl.GetFPS()

	if rl.GetMouseWheelMove() != 0 {
		camera.Zoom += rl.GetMouseWheelMove() / 10 * camera.Zoom
		if camera.Zoom < minZoom {
			camera.Zoom = minZoom
		} else if camera.Zoom > maxZoom {
			camera.Zoom = maxZoom
		}
		newX := player.Position.X + 16*global.Scale - global.ScreenWidth*global.Scale/camera.Zoom/2
		newY := player.Position.Y + 16*global.Scale - global.ScreenHeight*global.Scale/camera.Zoom/2

		movement.CorrectForZoom(newX, newY, camera)
	}

	if player.AttackOngoing {
		attack.Attack(player, fps)
	} else if time.Since(player.DashLastUse) < time.Millisecond*200 {
		movement.ContinueDash(player, camera, collisionObjects)
	} else {
		movement.Move(player, camera, fps, collisionObjects)

		if rl.IsKeyPressed(rl.KeySpace) &&
			player.DashCooldown.Milliseconds() < time.Since(player.DashLastUse).Milliseconds() {
			movement.Dash(player, camera, fps)
		}
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) &&
			player.AttackCooldown.Milliseconds() < time.Since(player.AttackLastUse).Milliseconds() {
			attack.StartAttack(player, bloonObjects)
		}
	}

	var objectsToRemove []int
	for i, object := range itemObjects {
		if object == nil {
			continue
		}
		if rl.CheckCollisionRecs(player.HitBox, object.HitBox) {
			objectsToRemove = append(objectsToRemove, i)
		}
	}

	for _, remove := range objectsToRemove {
		itemObjects = slices.Delete(itemObjects, remove, remove+1)
	}
}
