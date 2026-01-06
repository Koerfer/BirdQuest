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
	} else if player.Talking {
		performDialog(player)
	} else {
		door = movement.Move(player, camera)

		if rl.IsKeyPressed(rl.KeyE) {
			scene.AttemptQuestStep(player)
			player.IsMoving = false
			player.BaseRectangle = player.Animation.GetRectangleAreaInTexture(0)
			player.AnimationStep = 0
			player.Rotation = 0
		}

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

func performDialog(player *models.Player) {
	if !rl.IsKeyPressed(rl.KeySpace) {
		return
	}

	player.DialogStep++
	if player.DialogStep >= len(player.CurrentQuest.Steps[player.CurrentQuest.CurrentStep].Dialogs) {
		player.DialogStep = 0
		player.Talking = false
		player.DialogNPC = nil
		player.CurrentQuest.CurrentStep++
		if player.CurrentQuest.CurrentStep >= len(player.CurrentQuest.Steps) {
			player.CurrentQuest.Completed = true
			player.CurrentQuest = nil
		}
	}
}
