package models

import (
	"BirdQuest/global"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CollisionItems struct {
	DrawFirst   []*Object
	DrawDynamic []*Object
	DrawLast    []*Object
}

func (items *CollisionItems) DrawFirstLayer() {
	for _, item := range items.DrawFirst {
		if item == nil {
			continue
		}

		rl.DrawTexturePro(
			global.VariableSet.Textures32x32,
			*item.BaseRectangle,
			rl.Rectangle{
				X:      item.BasePositionRectangle.X * global.VariableSet.EntityScale,
				Y:      item.BasePositionRectangle.Y * global.VariableSet.EntityScale,
				Width:  item.BasePositionRectangle.Width * global.VariableSet.EntityScale,
				Height: item.BasePositionRectangle.Height * global.VariableSet.EntityScale,
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

func (items *CollisionItems) DrawDynamicLayer(player *Player) {
	var drawAfterPlayer []*Object

	for _, item := range items.DrawDynamic {
		if item == nil {
			continue
		}

		if item.BasePositionRectangle.Y >= player.BasePositionRectangle.Y {
			if drawAfterPlayer == nil {
				drawAfterPlayer = make([]*Object, 0)
			}

			drawAfterPlayer = append(drawAfterPlayer, item)
			continue
		}

		rl.DrawTexturePro(
			global.VariableSet.Textures32x32,
			*item.BaseRectangle,
			rl.Rectangle{
				X:      item.BasePositionRectangle.X * global.VariableSet.EntityScale,
				Y:      item.BasePositionRectangle.Y * global.VariableSet.EntityScale,
				Width:  item.BasePositionRectangle.Width * global.VariableSet.EntityScale,
				Height: item.BasePositionRectangle.Height * global.VariableSet.EntityScale,
			},
			rl.Vector2{
				X: 0,
				Y: 0,
			},
			0,
			rl.White,
		)
	}

	player.Draw()

	for _, item := range drawAfterPlayer {
		if item == nil {
			continue
		}

		rl.DrawTexturePro(
			global.VariableSet.Textures32x32,
			*item.BaseRectangle,
			rl.Rectangle{
				X:      item.BasePositionRectangle.X * global.VariableSet.EntityScale,
				Y:      item.BasePositionRectangle.Y * global.VariableSet.EntityScale,
				Width:  item.BasePositionRectangle.Width * global.VariableSet.EntityScale,
				Height: item.BasePositionRectangle.Height * global.VariableSet.EntityScale,
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

func (items *CollisionItems) DrawLastLayer() {
	for _, item := range items.DrawLast {
		if item == nil {
			continue
		}

		rl.DrawTexturePro(
			global.VariableSet.Textures32x32,
			*item.BaseRectangle,
			rl.Rectangle{
				X:      item.BasePositionRectangle.X * global.VariableSet.EntityScale,
				Y:      item.BasePositionRectangle.Y * global.VariableSet.EntityScale,
				Width:  item.BasePositionRectangle.Width * global.VariableSet.EntityScale,
				Height: item.BasePositionRectangle.Height * global.VariableSet.EntityScale,
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
