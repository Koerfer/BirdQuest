package initiate

import (
	"BirdQuest/global"
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
	"time"
)

func PreparePlayer() *models.Player {
	chiliAnimations := &models.Sprites{
		Texture:      global.VariableSet.Textures32x32,
		TileWidth:    global.TileWidth,
		TileHeight:   global.TileHeight,
		WidthInTiles: int(global.VariableSet.Textures32x32.Width) / global.TileWidth,
	}

	player := &models.Player{
		IsMoving:       false,
		AnimationStep:  0,
		Rotation:       0,
		Animation:      chiliAnimations,
		DashLastUse:    time.Time{},
		DashCooldown:   time.Millisecond * 1200,
		AttackLastUse:  time.Time{},
		AttackCooldown: time.Millisecond * 500,

		Object: models.Object{
			BaseRectangle: chiliAnimations.GetRectangleAreaInTexture(7),
		}}

	player.BasePositionRectangle = &rl.Rectangle{
		X:      0,
		Y:      0,
		Width:  player.BaseRectangle.Width,
		Height: player.BaseRectangle.Height,
	}

	return player
}
