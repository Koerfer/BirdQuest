package scene

import (
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func AttemptQuestStep(player *models.Player) {
	for _, quest := range Quests {
		if quest == nil {
			continue
		}

		if quest.Completed {
			continue
		}

		if quest.CurrentStep >= len(quest.Steps) {
			continue
		}

		switch quest.Steps[quest.CurrentStep].Type {
		case models.QuestStepInvalid:
			continue
		case models.QuestStepTalk:
			if rl.CheckCollisionRecs(createExtendedPlayerRectangle(player), *quest.Steps[quest.CurrentStep].NPC.BasePositionRectangle) {
				quest.Started = true
				player.Talking = true
				player.DialogNPC = quest.Steps[quest.CurrentStep].NPC
				player.CurrentQuest = quest
			}
		case models.QuestStepKill:

		case models.QuestStepOpenBox:
			if rl.CheckCollisionRecs(createExtendedPlayerRectangle(player), *quest.Steps[quest.CurrentStep].Box.BasePositionRectangle) {
				quest.Started = true
				player.CurrentQuest = quest

				box := quest.Steps[quest.CurrentStep].Box
				if box.OpeningStage < 2 {
					box.BaseRectangle.X += box.BaseRectangle.Width
					box.OpeningStage++
				} else if box.OpeningStage == 2 {
					box.BaseRectangle.X -= box.BaseRectangle.Width * 3
					box.BaseRectangle.Y += box.BaseRectangle.Height
					box.OpeningStage++
				} else if box.OpeningStage < 4 {
					box.BaseRectangle.X += box.BaseRectangle.Width
					box.OpeningStage++
				} else {
					player.CurrentQuest.CurrentStep++
					if player.CurrentQuest.CurrentStep >= len(player.CurrentQuest.Steps) {
						player.CurrentQuest.Completed = true
						player.CurrentQuest = nil
					}
				}
			}
		case models.QuestStepComplete:

		}
	}
}

func createExtendedPlayerRectangle(player *models.Player) rl.Rectangle {
	return rl.Rectangle{
		X:      player.BasePositionRectangle.X - 16,
		Y:      player.BasePositionRectangle.Y - 16,
		Width:  player.BasePositionRectangle.Width + 16,
		Height: player.BasePositionRectangle.Height + 16,
	}
}
