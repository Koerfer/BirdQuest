package models

import (
	"BirdQuest/global"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type NPC struct {
	Name          string
	StartedQuests []int

	Object
}

func (npc *NPC) Draw() {
	if npc == nil {
		return
	}

	rl.DrawTexturePro(
		global.VariableSet.Textures32x32,
		*npc.BaseRectangle,
		rl.Rectangle{
			X:      npc.BasePositionRectangle.X * global.VariableSet.EntityScale,
			Y:      npc.BasePositionRectangle.Y * global.VariableSet.EntityScale,
			Width:  npc.BasePositionRectangle.Width * global.VariableSet.EntityScale,
			Height: npc.BasePositionRectangle.Height * global.VariableSet.EntityScale,
		},
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)
}
