package update

import (
	"BirdQuest/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
	"slices"
)

func killItems(player *objects.Player, itemObjects []*objects.Object, bloonObjects []*objects.Bloon) {
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

	for i, bloon := range bloonObjects {
		if bloon == nil {
			continue
		}

		if bloon.PoppingAnimationStage == 0 {
			continue
		}

		bloon.AnimationStep++
		if bloon.AnimationStep == int(70*60/rl.GetFPS()) {
			bloon.AnimationStep = 0
			bloon.PoppingAnimationStage++
		}

		if bloon.PoppingAnimationStage == 5 {
			bloonObjects = slices.Delete(bloonObjects, i, i+1)
		}

	}
}
