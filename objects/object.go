package objects

import rl "github.com/gen2brain/raylib-go/raylib"

type Object struct {
	Position          rl.Vector2
	Texture           rl.Texture2D
	Rectangle         rl.Rectangle
	HitBox            rl.Rectangle
	AlwaysRenderLast  bool
	AlwaysRenderFirst bool
}

type Sprites struct {
	Texture      rl.Texture2D
	TileWidth    int
	TileHeight   int
	WidthInTiles int
	Properties   []*Property
}

type Property struct {
	id                int
	HitBoxX           int
	HitBoxY           int
	HitBoxWidth       int
	HitBoxHeight      int
	AlwaysRenderLast  bool
	AlwaysRenderFirst bool
}

func (s *Sprites) GetSrc(id int) rl.Rectangle {
	x := id % s.WidthInTiles
	y := id / s.WidthInTiles

	return rl.NewRectangle(
		float32(x*s.TileWidth),
		float32(y*s.TileHeight),
		float32(s.TileWidth),
		float32(s.TileHeight),
	)
}
