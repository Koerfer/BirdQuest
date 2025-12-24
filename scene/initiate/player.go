package initiate

import (
	"BirdQuest/global"
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
	"path/filepath"
	"time"
)

func PreparePlayer(path string) *models.Player {
	chiliAnimationsRaw := rl.LoadTexture(filepath.Join(path, "chili.png"))

	chiliAnimations := &models.Sprites{
		Texture:      chiliAnimationsRaw,
		TileWidth:    global.TileWidth,
		TileHeight:   global.TileHeight,
		WidthInTiles: int(chiliAnimationsRaw.Width) / global.TileWidth,
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

		Texture: chiliAnimations.Texture,
		Object: models.Object{
			BasePosition:  &rl.Vector2{},
			BaseRectangle: chiliAnimations.GetRectangleAreaInTexture(7),
		}}

	player.Rectangle = &rl.Rectangle{
		X:      player.BasePosition.X * global.VariableSet.EntityScale,
		Y:      player.BasePosition.Y * global.VariableSet.EntityScale,
		Width:  player.BaseRectangle.Width * global.VariableSet.EntityScale,
		Height: player.BaseRectangle.Height * global.VariableSet.EntityScale,
	}

	return player
}
