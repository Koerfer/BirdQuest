package update

import (
	"BirdQuest/attack"
	"BirdQuest/movement"
	"BirdQuest/scene"
	rl "github.com/gen2brain/raylib-go/raylib"
	"time"
)

func updatePlayer(camera *rl.Camera2D, player *scene.Player) {
	var door *scene.Door

	if player.AttackOngoing {
		attack.Attack(player)
	} else if time.Since(player.DashLastUse) < time.Millisecond*200 {
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
		scene.ChangeScene(door, player)
		movement.InitialiseCamera(player, camera)
	}
}
