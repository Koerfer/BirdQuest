package attack

import (
	"BirdQuest/global"
	"BirdQuest/scene"
	rl "github.com/gen2brain/raylib-go/raylib"
	"time"
)

var frameCounter int

func StartAttack(player *scene.Player) {
	frameCounter = 0

	player.AttackLastUse = time.Now()
	player.AttackOngoing = true

	player.AnimationStep = 3
	player.Rectangle = player.Animation.GetSrc(player.AnimationStep)

	var hit bool
	var hitId int
	extendedPlayerHitBox := &rl.Rectangle{
		X:      player.HitBox.X,
		Y:      player.HitBox.Y,
		Width:  player.HitBox.Width + 5*global.VariableSet.EntityScale,
		Height: player.HitBox.Height + 5*global.VariableSet.EntityScale,
	}
	for i, bloon := range scene.CurrentScene.BloonObjects {
		if bloon == nil {
			continue
		}
		if rl.CheckCollisionRecs(*extendedPlayerHitBox, bloon.HitBox) {
			hit = true
			hitId = i
			break
		}
	}
	extendedPlayerHitBox = nil

	if hit {
		bloons := scene.CurrentScene.BloonObjects
		if bloons[hitId].Lives == 0 {
			return
		}

		bloons[hitId].Lives--
		if bloons[hitId].Lives == 0 {
			bloons[hitId].PoppingAnimationStage = 1
		} else {
			bloons[hitId].Rectangle.X = float32(int(bloons[hitId].Rectangle.X+global.TileWidth) % int(bloons[hitId].Texture.Width))
			if bloons[hitId].Rectangle.X == 0 {
				bloons[hitId].Rectangle.X += global.TileWidth
			}
		}

		return
	}
}

func Attack(player *scene.Player) {

	if frameCounter >= int(4/global.VariableSet.FpsScale) {
		frameCounter = 0
		player.AnimationStep++
		if player.AnimationStep == 7 {
			player.AttackOngoing = false
			return
		}
		player.Rectangle = player.Animation.GetSrc(player.AnimationStep)
	}
	frameCounter++

}
