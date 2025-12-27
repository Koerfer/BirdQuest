package update

import (
	"BirdQuest/attack"
	"BirdQuest/global"
	"BirdQuest/movement"
	"BirdQuest/scene"
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
	"time"
)

func updatePlayer(camera *rl.Camera2D, player *models.Player) {
	var door *models.Door

	if player.AttackOngoing {
		attack.Attack(player)
	} else if time.Since(player.DashLastUse) < time.Millisecond*150 {
		door = movement.ContinueDash(player, camera)
	} else {
		door = movement.Move(player, camera)

		if rl.IsKeyPressed(rl.KeySpace) &&
			player.DashCooldown.Milliseconds() < time.Since(player.DashLastUse).Milliseconds() {
			movement.Dash(player, camera)
		}
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) &&
			player.AttackCooldown.Milliseconds() < time.Since(player.AttackLastUse).Milliseconds() {
			attack.StartAttack(player)
		}
	}

	if door != nil {
		previousHeight := scene.CurrentScene.Height
		scene.ChangeScene(door, player)

		zoom := scene.CurrentScene.Height * camera.Zoom / previousHeight
		if zoom < 1 {
			zoom = 1
		} else if zoom > maxZoom {
			zoom = maxZoom
		}

		global.Zoom(zoom, camera)

		movement.InitialiseCamera(player, camera)
	}
}
