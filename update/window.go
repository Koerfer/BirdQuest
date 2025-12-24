package update

import (
	"BirdQuest/global"
	"BirdQuest/movement"
	"BirdQuest/scene"
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Window(player *models.Player, camera *rl.Camera2D) {
	if rl.IsKeyPressed(rl.KeyF5) && !rl.IsWindowFullscreen() {
		rl.ToggleFullscreen()
		updateDesiredWindowSize(float32(rl.GetMonitorWidth(rl.GetCurrentMonitor())), float32(rl.GetMonitorHeight(rl.GetCurrentMonitor())), player, camera)
	} else if rl.IsKeyPressed(rl.KeyF5) && rl.IsWindowFullscreen() {
		rl.ToggleFullscreen()
		updateDesiredWindowSize(float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight()), player, camera)
	}

	if rl.IsWindowResized() {
		updateDesiredWindowSize(float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight()), player, camera)
	}
}

func updateDesiredWindowSize(width, height float32, player *models.Player, camera *rl.Camera2D) {
	player.BasePosition.X = player.Rectangle.X / global.VariableSet.EntityScale
	player.BasePosition.Y = player.Rectangle.Y / global.VariableSet.EntityScale

	global.VariableSet.DesiredWidth = width
	global.VariableSet.DesiredHeight = height

	global.VariableSet.EntityScale = global.VariableSet.DesiredWidth / scene.CurrentScene.Width
	global.VariableSet.ScaleHeight = global.VariableSet.DesiredHeight / scene.CurrentScene.Height
	global.VariableSet.ScaleWidth = global.VariableSet.DesiredWidth / scene.CurrentScene.Width

	if global.VariableSet.EntityScale < global.VariableSet.DesiredHeight/scene.CurrentScene.Height {
		global.VariableSet.EntityScale = global.VariableSet.DesiredHeight / scene.CurrentScene.Height
	}

	global.VariableSet.MapHeight = scene.CurrentScene.Height * global.VariableSet.EntityScale
	global.VariableSet.MapWidth = scene.CurrentScene.Width * global.VariableSet.EntityScale

	global.VariableSet.PlayerMiddleOffset = global.TileWidth / 2 * global.VariableSet.EntityScale
	global.VariableSet.EntitySize = global.TileWidth * global.VariableSet.EntityScale

	global.VariableSet.Speed = global.VariableSet.FpsScale * global.VariableSet.EntityScale

	global.VariableSet.VisibleMapHeight = global.VariableSet.DesiredHeight / camera.Zoom
	global.VariableSet.VisibleMapWidth = global.VariableSet.DesiredWidth / camera.Zoom

	for _, item := range scene.CurrentScene.ItemObjects.Objects {
		item.Rectangle = &rl.Rectangle{
			X:      item.BasePosition.X * global.VariableSet.EntityScale,
			Y:      item.BasePosition.Y * global.VariableSet.EntityScale,
			Width:  item.BaseRectangle.Width * global.VariableSet.EntityScale,
			Height: item.BaseRectangle.Height * global.VariableSet.EntityScale,
		}
	}

	for _, bloon := range scene.CurrentScene.Bloons.BloonObjects {
		bloon.Rectangle = &rl.Rectangle{
			X:      bloon.BasePosition.X * global.VariableSet.EntityScale,
			Y:      bloon.BasePosition.Y * global.VariableSet.EntityScale,
			Width:  bloon.BaseRectangle.Width * global.VariableSet.EntityScale,
			Height: bloon.BaseRectangle.Height * global.VariableSet.EntityScale,
		}
	}

	for _, collisionItem := range scene.CurrentScene.CollisionObjects.DrawFirst {
		collisionItem.Rectangle = &rl.Rectangle{
			X:      collisionItem.BasePosition.X * global.VariableSet.EntityScale,
			Y:      collisionItem.BasePosition.Y * global.VariableSet.EntityScale,
			Width:  collisionItem.BaseRectangle.Width * global.VariableSet.EntityScale,
			Height: collisionItem.BaseRectangle.Height * global.VariableSet.EntityScale,
		}
	}

	for _, collisionItem := range scene.CurrentScene.CollisionObjects.DrawDynamic {
		collisionItem.Rectangle = &rl.Rectangle{
			X:      collisionItem.BasePosition.X * global.VariableSet.EntityScale,
			Y:      collisionItem.BasePosition.Y * global.VariableSet.EntityScale,
			Width:  collisionItem.BaseRectangle.Width * global.VariableSet.EntityScale,
			Height: collisionItem.BaseRectangle.Height * global.VariableSet.EntityScale,
		}
	}

	for _, collisionItem := range scene.CurrentScene.CollisionObjects.DrawLast {
		collisionItem.Rectangle = &rl.Rectangle{
			X:      collisionItem.BasePosition.X * global.VariableSet.EntityScale,
			Y:      collisionItem.BasePosition.Y * global.VariableSet.EntityScale,
			Width:  collisionItem.BaseRectangle.Width * global.VariableSet.EntityScale,
			Height: collisionItem.BaseRectangle.Height * global.VariableSet.EntityScale,
		}

	}

	for i, collisionBox := range scene.CurrentScene.CollisionBoxes {
		collisionBox.X = scene.CurrentScene.BaseCollisionBoxes[i].X * global.VariableSet.EntityScale
		collisionBox.Y = scene.CurrentScene.BaseCollisionBoxes[i].Y * global.VariableSet.EntityScale
		collisionBox.Width = scene.CurrentScene.BaseCollisionBoxes[i].Width * global.VariableSet.EntityScale
		collisionBox.Height = scene.CurrentScene.BaseCollisionBoxes[i].Height * global.VariableSet.EntityScale
	}

	for _, door := range scene.CurrentScene.Doors {
		door.Rectangle = &rl.Rectangle{
			X:      door.BaseRectangle.X * global.VariableSet.EntityScale,
			Y:      door.BaseRectangle.Y * global.VariableSet.EntityScale,
			Width:  door.BaseRectangle.Width * global.VariableSet.EntityScale,
			Height: door.BaseRectangle.Height * global.VariableSet.EntityScale,
		}
	}

	player.Rectangle.X = player.BasePosition.X * global.VariableSet.EntityScale
	player.Rectangle.Y = player.BasePosition.Y * global.VariableSet.EntityScale
	player.Rectangle.Width = player.BaseRectangle.Width * global.VariableSet.EntityScale
	player.Rectangle.Height = player.BaseRectangle.Height * global.VariableSet.EntityScale

	movement.InitialiseCamera(player, camera)
}
