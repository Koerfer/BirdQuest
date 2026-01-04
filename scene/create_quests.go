package scene

import (
	"BirdQuest/scene/models"
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

	stepOne := &models.Step{
		Type: models.QuestStepTalk,
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

	for _, npc := range AllScenes["main"].NPCs {
		if npc.Name == "BlueTit" {
			stepOne.NPC = npc
		}
	}

	quest.Steps = append(quest.Steps, stepOne)

	stepTwo := &models.Step{
		Type: models.QuestStepOpenBox,
		Box:  AllScenes["main"].SeedBoxes[0],
	}

	quest.Steps = append(quest.Steps, stepTwo)
	Quests = append(Quests, quest)
}
