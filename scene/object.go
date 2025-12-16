package scene

import rl "github.com/gen2brain/raylib-go/raylib"

type Items struct {
	Objects []*Object
	Texture rl.Texture2D
}

type CollisionItems struct {
	DrawFirst   []*Object
	DrawDynamic []*Object
	DrawLast    []*Object

	Texture rl.Texture2D
}

type Object struct {
	Position  rl.Vector2
	Rectangle rl.Rectangle
	HitBox    rl.Rectangle
}

type sprites struct {
	Texture      rl.Texture2D
	TileWidth    int
	TileHeight   int
	WidthInTiles int
	Properties   []*property
}

type property struct {
	id                int
	HitBoxX           int
	HitBoxY           int
	HitBoxWidth       int
	HitBoxHeight      int
	AlwaysRenderLast  bool
	AlwaysRenderFirst bool
}

func (s *sprites) GetRectangleAreaInTexture(id int) rl.Rectangle {
	x := id % s.WidthInTiles
	y := id / s.WidthInTiles

	return rl.NewRectangle(
		float32(x*s.TileWidth),
		float32(y*s.TileHeight),
		float32(s.TileWidth),
		float32(s.TileHeight),
	)
}
