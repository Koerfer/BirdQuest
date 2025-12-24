package models

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CollisionItems struct {
	DrawFirst   []*Object
	DrawDynamic []*Object
	DrawLast    []*Object

	Texture rl.Texture2D
}

func (items *CollisionItems) DrawFirstLayer() {
	for _, item := range items.DrawFirst {
		if item == nil {
			continue
		}

		rl.DrawTexturePro(
			items.Texture,
			*item.BaseRectangle,
			rl.Rectangle{
				X:      item.Rectangle.X,
				Y:      item.Rectangle.Y,
				Width:  item.Rectangle.Width,
				Height: item.Rectangle.Height,
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

		if item.Rectangle.Y >= player.Rectangle.Y {
			if drawAfterPlayer == nil {
				drawAfterPlayer = make([]*Object, 0)
			}

			drawAfterPlayer = append(drawAfterPlayer, item)
			continue
		}

		rl.DrawTexturePro(
			items.Texture,
			*item.BaseRectangle,
			rl.Rectangle{
				X:      item.Rectangle.X,
				Y:      item.Rectangle.Y,
				Width:  item.Rectangle.Width,
				Height: item.Rectangle.Height,
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
			items.Texture,
			*item.BaseRectangle,
			rl.Rectangle{
				X:      item.Rectangle.X,
				Y:      item.Rectangle.Y,
				Width:  item.Rectangle.Width,
				Height: item.Rectangle.Height,
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
			items.Texture,
			*item.BaseRectangle,
			rl.Rectangle{
				X:      item.Rectangle.X,
				Y:      item.Rectangle.Y,
				Width:  item.Rectangle.Width,
				Height: item.Rectangle.Height,
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
