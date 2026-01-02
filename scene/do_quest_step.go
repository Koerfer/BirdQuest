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
			extendedPlayerRectangle := rl.Rectangle{
				X:      player.BasePositionRectangle.X - 16,
				Y:      player.BasePositionRectangle.Y - 16,
				Width:  player.BasePositionRectangle.Width + 16,
				Height: player.BasePositionRectangle.Height + 16,
			}
			if rl.CheckCollisionRecs(extendedPlayerRectangle, *quest.Steps[quest.CurrentStep].NPC.BasePositionRectangle) {
				quest.Started = true
				player.Talking = true
				player.DialogNPC = quest.Steps[quest.CurrentStep].NPC
				player.CurrentQuest = quest
			}
		case models.QuestStepKill:

		case models.QuestStepOpenBox:

		case models.QuestStepComplete:

		}
	}
}
