package update

import (
	"BirdQuest/attack"
	"BirdQuest/movement"
	"BirdQuest/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
	"time"
)

func updatePlayer(camera *rl.Camera2D, player *objects.Player, collisionObjects []*rl.Rectangle, bloonObjects []*objects.Bloon) {
	if player.AttackOngoing {
		attack.Attack(player)
	} else if time.Since(player.DashLastUse) < time.Millisecond*200 {
		movement.ContinueDash(player, camera, collisionObjects)
	} else {
		movement.Move(player, camera, collisionObjects)

		if rl.IsKeyPressed(rl.KeySpace) &&
			player.DashCooldown.Milliseconds() < time.Since(player.DashLastUse).Milliseconds() {
			movement.Dash(player, camera)
		}
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) &&
			player.AttackCooldown.Milliseconds() < time.Since(player.AttackLastUse).Milliseconds() {
			attack.StartAttack(player, bloonObjects)
		}
	}
}
