package models

import rl "github.com/gen2brain/raylib-go/raylib"

type Sprites struct {
	Texture      rl.Texture2D
	TileWidth    int
	TileHeight   int
	WidthInTiles int
	Properties   []*Property
}

type Property struct {
	Id                int
	AlwaysRenderLast  bool
	AlwaysRenderFirst bool
}

func (s *Sprites) GetRectangleAreaInTexture(id int) *rl.Rectangle {
	x := id % s.WidthInTiles
	y := id / s.WidthInTiles

	return &rl.Rectangle{
		X:      float32(x * s.TileWidth),
		Y:      float32(y * s.TileHeight),
		Width:  float32(s.TileWidth),
		Height: float32(s.TileHeight),
	}
}
