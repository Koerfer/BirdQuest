package scene

import (
	"BirdQuest/global"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (items *Items) Draw() {
	for _, item := range items.Objects {
		if item == nil {
			continue
		}

		rl.DrawTexturePro(
			items.Texture,
			item.Rectangle,
			rl.Rectangle{
				X:      item.Position.X,
				Y:      item.Position.Y,
				Width:  item.Rectangle.Width * global.VariableSet.EntityScale,
				Height: item.Rectangle.Height * global.VariableSet.EntityScale,
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

func (bloons *Bloons) Draw() {

	for _, bloon := range bloons.BloonObjects {
		if bloon == nil {
			continue
		}

		// If bloon is popped
		if bloon.PoppingAnimationStage > 0 {
			bloon.Rectangle = rl.Rectangle{
				X:      global.TileWidth * (float32(bloon.PoppingAnimationStage + 2)),
				Y:      0,
				Width:  bloon.Rectangle.Width,
				Height: bloon.Rectangle.Height,
			}
		}

		rl.DrawTexturePro(
			bloons.Texture,
			bloon.Rectangle,
			rl.Rectangle{
				X:      bloon.Position.X,
				Y:      bloon.Position.Y,
				Width:  bloon.Rectangle.Width * global.VariableSet.EntityScale,
				Height: bloon.Rectangle.Height * global.VariableSet.EntityScale,
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

func (items *CollisionItems) DrawFirstLayer() {
	for _, item := range items.DrawFirst {
		if item == nil {
			continue
		}

		rl.DrawTexturePro(
			items.Texture,
			item.Rectangle,
			rl.Rectangle{
				X:      item.Position.X,
				Y:      item.Position.Y,
				Width:  item.Rectangle.Width * global.VariableSet.EntityScale,
				Height: item.Rectangle.Height * global.VariableSet.EntityScale,
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

		if item.Position.Y >= player.Position.Y {
			if drawAfterPlayer == nil {
				drawAfterPlayer = make([]*Object, 0)
			}

			drawAfterPlayer = append(drawAfterPlayer, item)
			continue
		}

		rl.DrawTexturePro(
			items.Texture,
			item.Rectangle,
			rl.Rectangle{
				X:      item.Position.X,
				Y:      item.Position.Y,
				Width:  item.Rectangle.Width * global.VariableSet.EntityScale,
				Height: item.Rectangle.Height * global.VariableSet.EntityScale,
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
			item.Rectangle,
			rl.Rectangle{
				X:      item.Position.X,
				Y:      item.Position.Y,
				Width:  item.Rectangle.Width * global.VariableSet.EntityScale,
				Height: item.Rectangle.Height * global.VariableSet.EntityScale,
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
			item.Rectangle,
			rl.Rectangle{
				X:      item.Position.X,
				Y:      item.Position.Y,
				Width:  item.Rectangle.Width * global.VariableSet.EntityScale,
				Height: item.Rectangle.Height * global.VariableSet.EntityScale,
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
func (player *Player) Draw() {
	var shadowOffset float32 = 20
	if !player.IsMoving && !player.AttackOngoing {
		shadowOffset = 2
	}

	rl.DrawTexturePro(
		player.Texture,
		rl.Rectangle{
			X:      player.Rectangle.X,
			Y:      player.Rectangle.Y + 96,
			Width:  player.Rectangle.Width,
			Height: player.Rectangle.Height,
		},
		rl.Rectangle{
			X:      player.Position.X + global.VariableSet.PlayerMiddleOffset,
			Y:      player.Position.Y + global.VariableSet.PlayerMiddleOffset + shadowOffset*global.VariableSet.EntityScale,
			Width:  player.Rectangle.Width * global.VariableSet.EntityScale,
			Height: player.Rectangle.Height * global.VariableSet.EntityScale,
		},
		rl.Vector2{X: global.VariableSet.PlayerMiddleOffset, Y: global.VariableSet.PlayerMiddleOffset},
		player.Rotation,
		rl.White,
	)

	rl.DrawTexturePro(
		player.Texture,
		player.Rectangle,
		rl.Rectangle{
			X:      player.Position.X + global.VariableSet.PlayerMiddleOffset,
			Y:      player.Position.Y + global.VariableSet.PlayerMiddleOffset,
			Width:  player.Rectangle.Width * global.VariableSet.EntityScale,
			Height: player.Rectangle.Height * global.VariableSet.EntityScale,
		},
		rl.Vector2{X: global.VariableSet.PlayerMiddleOffset, Y: global.VariableSet.PlayerMiddleOffset},
		player.Rotation,
		rl.White,
	)
}
