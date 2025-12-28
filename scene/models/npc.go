package models

import (
	"BirdQuest/global"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type NPC struct {
	Name          string
	StartedQuests []int
	BaseRectangle *rl.Rectangle
}

func (npc *NPC) Draw() {
	if npc == nil {
		return
	}

	rl.DrawRectangleRec(
		rl.Rectangle{
			X:      npc.BaseRectangle.X * global.VariableSet.EntityScale,
			Y:      npc.BaseRectangle.Y * global.VariableSet.EntityScale,
			Width:  npc.BaseRectangle.Width * global.VariableSet.EntityScale,
			Height: npc.BaseRectangle.Height * global.VariableSet.EntityScale,
		},
		rl.Pink,
	)
}
