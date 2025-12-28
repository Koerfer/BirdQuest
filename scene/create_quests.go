package scene

import (
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var Quests []*models.Quest

func CreateQuests() {
	Quests = make([]*models.Quest, 0)
	quest := &models.Quest{
		Started:      false,
		Completed:    false,
		Dependencies: nil,
		Steps:        make([]*models.Step, 0),
		CurrentStep:  0,
	}

	blueTit := &models.NPC{
		Name: "Blue Tit",
		BaseRectangle: &rl.Rectangle{
			X:      64,
			Y:      32,
			Width:  32,
			Height: 32,
		},
	}

	stepOne := &models.Step{
		Type: models.QuestStepTalk,
		NPC:  blueTit,
		Dialogs: []*models.Dialog{
			{
				NPCTalking: false,
				Line:       "Hello little bird!",
			}, {
				NPCTalking: true,
				Line:       "*chirp* Hi!",
			}, {
				NPCTalking: false,
				Line:       "What are you doing here?",
			}, {
				NPCTalking: true,
				Line:       "I'm trying to get this box opened",
			},
		},
	}

	quest.Steps = append(quest.Steps, stepOne)

	AllScenes["main"].NPCs = append(make([]*models.NPC, 0), blueTit)
	Quests = append(Quests, quest)
}
