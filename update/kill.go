package update

import (
	"BirdQuest/global"
	"BirdQuest/scene"
	rl "github.com/gen2brain/raylib-go/raylib"
	"slices"
)

func killItems(player *scene.Player) {
	var objectsToRemove []int
	for i, object := range scene.CurrentScene.ItemObjects.Objects {
		if object == nil {
			continue
		}
		if rl.CheckCollisionRecs(player.HitBox, object.HitBox) {
			objectsToRemove = append(objectsToRemove, i)
		}
	}

	for _, remove := range objectsToRemove {
		scene.CurrentScene.ItemObjects.Objects = slices.Delete(scene.CurrentScene.ItemObjects.Objects, remove, remove+1)
	}

	for i, bloon := range scene.CurrentScene.Bloons.BloonObjects {
		if bloon == nil {
			continue
		}

		if bloon.PoppingAnimationStage == 0 {
			continue
		}

		bloon.AnimationStep++
		if bloon.AnimationStep == int(35/global.VariableSet.FpsScale) {
			bloon.AnimationStep = 0
			bloon.PoppingAnimationStage++
		}

		if bloon.PoppingAnimationStage == 5 {
			scene.CurrentScene.Bloons.BloonObjects = slices.Delete(scene.CurrentScene.Bloons.BloonObjects, i, i+1)
		}

	}
}
