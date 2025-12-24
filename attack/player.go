package attack

import (
	"BirdQuest/global"
	"BirdQuest/scene"
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
	"time"
)

var frameCounter int

func StartAttack(player *models.Player) {
	frameCounter = 0

	player.AttackLastUse = time.Now()
	player.AttackOngoing = true

	player.AnimationStep = 3
	player.BaseRectangle = player.Animation.GetRectangleAreaInTexture(player.AnimationStep)

	var hit bool
	var hitId int
	extendedPlayerRectangle := &rl.Rectangle{
		X:      player.Rectangle.X,
		Y:      player.Rectangle.Y,
		Width:  player.Rectangle.Width + 5*global.VariableSet.EntityScale,
		Height: player.Rectangle.Height + 5*global.VariableSet.EntityScale,
	}
	for i, bloon := range scene.CurrentScene.Bloons.BloonObjects {
		if bloon == nil {
			continue
		}
		if rl.CheckCollisionRecs(*extendedPlayerRectangle, *bloon.Rectangle) {
			hit = true
			hitId = i
			break
		}
	}
	extendedPlayerRectangle = nil

	if hit {
		bloons := scene.CurrentScene.Bloons.BloonObjects
		if bloons[hitId].Lives == 0 {
			return
		}

		bloons[hitId].Lives--
		if bloons[hitId].Lives == 0 {
			bloons[hitId].PoppingAnimationStage = 1
		} else {
			bloons[hitId].BaseRectangle.X = float32(int(bloons[hitId].BaseRectangle.X+global.TileWidth) % int(scene.CurrentScene.Bloons.Texture.Width))
			if bloons[hitId].BaseRectangle.X == 0 {
				bloons[hitId].BaseRectangle.X += global.TileWidth
			}
		}

		return
	}
}

func Attack(player *models.Player) {

	if frameCounter >= int(4/global.VariableSet.FpsScale) {
		frameCounter = 0
		player.AnimationStep++
		if player.AnimationStep == 7 {
			player.AttackOngoing = false
			return
		}
		player.BaseRectangle = player.Animation.GetRectangleAreaInTexture(player.AnimationStep)
	}
	frameCounter++

}
