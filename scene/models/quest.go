package models

import (
	"BirdQuest/global"
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

type Quest struct {
	Started      bool
	Completed    bool
	Dependencies []*Quest

	Steps       []*Step
	CurrentStep int
}

type Step struct {
	Type    QuestStep
	NPC     *NPC
	Box     *SeedBox
	Dialogs []*Dialog
}

type Dialog struct {
	NPCTalking bool
	Line       string
}

type QuestStep int8

const (
	QuestStepInvalid QuestStep = iota
	QuestStepTalk
	QuestStepKill
	QuestStepOpenBox
	QuestStepComplete
)

func (quest *Quest) Draw() {
	if quest.Completed {
		return
	}

	var questMarkerRectangle rl.Rectangle
	if !quest.Started {
		questMarkerRectangle = rl.Rectangle{
			X:      608,
			Y:      608,
			Width:  32,
			Height: 32,
		}
	} else {
		questMarkerRectangle = rl.Rectangle{
			X:      576,
			Y:      608,
			Width:  32,
			Height: 32,
		}
	}

	if quest.Steps[quest.CurrentStep].Box != nil {
		rl.DrawTexturePro(
			global.VariableSet.Textures32x32,
			questMarkerRectangle,
			rl.Rectangle{
				X:      quest.Steps[quest.CurrentStep].Box.BasePositionRectangle.X * global.VariableSet.EntityScale,
				Y:      (quest.Steps[quest.CurrentStep].Box.BasePositionRectangle.Y - 32) * global.VariableSet.EntityScale,
				Width:  32 * global.VariableSet.EntityScale,
				Height: 32 * global.VariableSet.EntityScale,
			},
			rl.Vector2{
				X: 0,
				Y: 0,
			},
			0,
			rl.White,
		)
	}
	if quest.Steps[quest.CurrentStep].NPC != nil {
		rl.DrawTexturePro(
			global.VariableSet.Textures32x32,
			questMarkerRectangle,
			rl.Rectangle{
				X:      quest.Steps[quest.CurrentStep].NPC.BasePositionRectangle.X * global.VariableSet.EntityScale,
				Y:      (quest.Steps[quest.CurrentStep].NPC.BasePositionRectangle.Y - 32) * global.VariableSet.EntityScale,
				Width:  32 * global.VariableSet.EntityScale,
				Height: 32 * global.VariableSet.EntityScale,
			},
			rl.Vector2{
				X: 0,
				Y: 0,
			},
			0,
			rl.White,
		)
	}
}

func (dialog *Dialog) Draw(camera rl.Camera2D, npc *NPC, player *Player) {
	dialogBackground := rl.Rectangle{
		X:      0,
		Y:      camera.Target.Y + global.VariableSet.VisibleMapHeight - 100*global.VariableSet.EntityScale/camera.Zoom,
		Width:  global.VariableSet.VisibleMapWidth,
		Height: 100 * global.VariableSet.EntityScale / camera.Zoom,
	}
	rl.DrawRectanglePro(
		dialogBackground,
		rl.Vector2{},
		0,
		color.RGBA{A: 200},
	)

	if !dialog.NPCTalking {
		rl.DrawTexturePro(
			global.VariableSet.Textures32x32,
			*player.BaseRectangle,
			rl.Rectangle{
				X:      dialogBackground.X + 50*global.VariableSet.EntityScale/camera.Zoom,
				Y:      dialogBackground.Y + 50*global.VariableSet.EntityScale/camera.Zoom,
				Width:  100 * global.VariableSet.EntityScale / camera.Zoom,
				Height: 100 * global.VariableSet.EntityScale / camera.Zoom,
			},
			rl.Vector2{
				X: 50 * global.VariableSet.EntityScale / camera.Zoom,
				Y: 50 * global.VariableSet.EntityScale / camera.Zoom,
			},
			player.Rotation,
			rl.White,
		)
	} else {
		rl.DrawTexturePro(
			global.VariableSet.Textures32x32,
			*npc.BaseRectangle,
			rl.Rectangle{
				X:      dialogBackground.X + 50*global.VariableSet.EntityScale/camera.Zoom,
				Y:      dialogBackground.Y + 50*global.VariableSet.EntityScale/camera.Zoom,
				Width:  100 * global.VariableSet.EntityScale / camera.Zoom,
				Height: 100 * global.VariableSet.EntityScale / camera.Zoom,
			},
			rl.Vector2{
				X: 50 * global.VariableSet.EntityScale / camera.Zoom,
				Y: 50 * global.VariableSet.EntityScale / camera.Zoom,
			},
			0,
			rl.White,
		)
	}

	textSize := rl.MeasureTextEx(
		global.Font,
		dialog.Line,
		32*global.VariableSet.EntityScale/camera.Zoom,
		2*global.VariableSet.EntityScale/camera.Zoom,
	)

	rl.DrawTextPro(global.Font,
		dialog.Line,
		rl.NewVector2(
			100*global.VariableSet.EntityScale/camera.Zoom+(dialogBackground.Width-100*global.VariableSet.EntityScale/camera.Zoom)/2,
			dialogBackground.Y+dialogBackground.Height/2,
		),
		rl.NewVector2(textSize.X/2, textSize.Y/2),
		0,
		32*global.VariableSet.EntityScale/camera.Zoom,
		2*global.VariableSet.EntityScale/camera.Zoom,
		rl.White,
	)
}
