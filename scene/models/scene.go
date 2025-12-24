package models

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct {
	Name               string
	Background         rl.Texture2D
	ItemObjects        *Items
	BaseCollisionBoxes []*rl.Rectangle
	CollisionBoxes     []*rl.Rectangle
	CollisionObjects   *CollisionItems
	Doors              []*Door
	Bloons             *Bloons

	Width  float32
	Height float32

	WidthInTiles  int
	HeightInTiles int
}
