package movement

import (
	"BirdQuest/global"
	"BirdQuest/scene"
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
	"time"
)

var frameCounter int

func Dash(player *models.Player, camera *rl.Camera2D) {
	player.DashLastUse = time.Now()

	mousePositionAbsolute := rl.GetMousePosition()
	mousePositionRelative := rl.Vector2{
		X: mousePositionAbsolute.X/camera.Zoom + camera.Target.X - (player.BasePositionRectangle.X*global.VariableSet.EntityScale + global.VariableSet.PlayerMiddleOffset),
		Y: mousePositionAbsolute.Y/camera.Zoom + camera.Target.Y - (player.BasePositionRectangle.Y*global.VariableSet.EntityScale + global.VariableSet.PlayerMiddleOffset),
	}

	dashDirection := rl.Vector2Normalize(mousePositionRelative)
	if math.Signbit(float64(dashDirection.X)) {
		player.Rotation = float32(360 - math.Acos(float64(-dashDirection.Y))*180/math.Pi)
	} else {
		player.Rotation = float32(math.Acos(float64(-dashDirection.Y)) * 180 / math.Pi)
	}

	dashDirection.X *= 20 * global.VariableSet.FpsScale
	dashDirection.Y *= 20 * global.VariableSet.FpsScale
	player.DashDirection = dashDirection
}

func ContinueDash(player *models.Player, camera *rl.Camera2D) *models.Door {
	player.IsMoving = true
	player.AnimationStep = 2
	player.BaseRectangle = player.Animation.GetRectangleAreaInTexture(player.AnimationStep)

	var goThroughDoor *models.Door

	if player.DashDirection.Y < 0 {
		goThroughDoor = moveUp(player, camera, -player.DashDirection.Y)
	} else {
		goThroughDoor = moveDown(player, camera, player.DashDirection.Y)
	}

	if player.DashDirection.X < 0 {
		goThroughDoor = moveLeft(player, camera, -player.DashDirection.X)
	} else {
		goThroughDoor = moveRight(player, camera, player.DashDirection.X)
	}

	return goThroughDoor
}

func Move(player *models.Player, camera *rl.Camera2D) *models.Door {
	up := rl.IsKeyDown(rl.KeyW)
	down := rl.IsKeyDown(rl.KeyS)
	left := rl.IsKeyDown(rl.KeyA)
	right := rl.IsKeyDown(rl.KeyD)

	diagonalSpeed := global.VariableSet.FpsScale * 3.535533
	normalSpeed := global.VariableSet.FpsScale * 5

	if player.BasePositionRectangle.Y*global.VariableSet.EntityScale < 0 {
		up = false
	}
	if player.BasePositionRectangle.Y*global.VariableSet.EntityScale+global.VariableSet.EntitySize > global.VariableSet.MapHeight {
		down = false
	}
	if player.BasePositionRectangle.X*global.VariableSet.EntityScale+global.VariableSet.EntitySize > global.VariableSet.MapWidth {
		right = false
	}
	if player.BasePositionRectangle.X < 0 {
		left = false
	}

	if down && up {
		player.IsMoving = false
		up = false
		down = false
	}
	if left && right {
		player.IsMoving = false
		left = false
		right = false
	}

	var goThroughDoor *models.Door

	if up && right {
		player.IsMoving = true
		player.Rotation = 45

		goThroughDoor = moveUp(player, camera, diagonalSpeed)
		goThroughDoor = moveRight(player, camera, diagonalSpeed)
	} else if up && left {
		player.IsMoving = true
		player.Rotation = 315

		goThroughDoor = moveUp(player, camera, diagonalSpeed)
		goThroughDoor = moveLeft(player, camera, diagonalSpeed)
	} else if down && left {
		player.IsMoving = true
		player.Rotation = 225

		goThroughDoor = moveDown(player, camera, diagonalSpeed)
		goThroughDoor = moveLeft(player, camera, diagonalSpeed)
	} else if down && right {
		player.IsMoving = true
		player.Rotation = 135

		goThroughDoor = moveDown(player, camera, diagonalSpeed)
		goThroughDoor = moveRight(player, camera, diagonalSpeed)
	} else if up {
		player.IsMoving = true
		player.Rotation = 0

		goThroughDoor = moveUp(player, camera, normalSpeed)
	} else if down {
		player.IsMoving = true
		player.Rotation = 180
		goThroughDoor = moveDown(player, camera, normalSpeed)
	} else if right {
		player.IsMoving = true
		player.Rotation = 90

		goThroughDoor = moveRight(player, camera, normalSpeed)
	} else if left {
		player.IsMoving = true
		player.Rotation = 270

		goThroughDoor = moveLeft(player, camera, normalSpeed)
	} else {
		player.IsMoving = false
	}

	if player.IsMoving {
		if frameCounter > int(8/global.VariableSet.FpsScale) {
			frameCounter = 0
			player.AnimationStep = (player.AnimationStep+1)%3 + 1
			player.BaseRectangle = player.Animation.GetRectangleAreaInTexture(player.AnimationStep)
		}

		frameCounter++
	}

	if !player.IsMoving {
		player.BaseRectangle = player.Animation.GetRectangleAreaInTexture(0)
		player.AnimationStep = 0
		player.Rotation = 0
		player.IsMoving = false
	}

	return goThroughDoor
}

func moveUp(player *models.Player, camera *rl.Camera2D, offset float32) *models.Door {
	var collided bool
	var collisionPoint float32
	var lastPosition float32

	newPlayerRectangleY := player.BasePositionRectangle.Y - offset
	for _, collisionBox := range scene.CurrentScene.BaseCollisionBoxes {
		if player.BasePositionRectangle.X+player.BasePositionRectangle.Width > collisionBox.X &&
			player.BasePositionRectangle.X < collisionBox.X+collisionBox.Width &&
			player.BasePositionRectangle.Y > collisionBox.Y &&
			newPlayerRectangleY <= collisionBox.Y+collisionBox.Height {

			collided = true
			lastPosition = player.BasePositionRectangle.Y

			if collisionBox.Y+collisionBox.Height > collisionPoint {
				collisionPoint = collisionBox.Y + collisionBox.Height
			}
		}
	}

	if collided {
		player.BasePositionRectangle.Y = collisionPoint
		moveCameraUp(player, camera, lastPosition-collisionPoint)
		return nil
	}

	for _, door := range scene.CurrentScene.Doors {
		if player.BasePositionRectangle.X+player.BasePositionRectangle.Width > door.BaseRectangle.X &&
			player.BasePositionRectangle.X < door.BaseRectangle.X+door.BaseRectangle.Width &&
			player.BasePositionRectangle.Y > door.BaseRectangle.Y &&
			newPlayerRectangleY < door.BaseRectangle.Y+door.BaseRectangle.Height {

			return door
		}
	}

	if player.BasePositionRectangle.Y-offset < 0 {
		player.BasePositionRectangle.Y = 0
	} else {
		player.BasePositionRectangle.Y -= offset
	}

	moveCameraUp(player, camera, offset)
	return nil
}

func moveDown(player *models.Player, camera *rl.Camera2D, offset float32) *models.Door {
	var collided bool
	var collisionPoint float32 = math.MaxFloat32
	var lastPosition float32

	newPlayerRectangleY := player.BasePositionRectangle.Y + offset
	for _, collisionBox := range scene.CurrentScene.BaseCollisionBoxes {
		if player.BasePositionRectangle.X+player.BasePositionRectangle.Width > collisionBox.X &&
			player.BasePositionRectangle.X < collisionBox.X+collisionBox.Width &&
			player.BasePositionRectangle.Y+player.BasePositionRectangle.Height <= collisionBox.Y &&
			newPlayerRectangleY+player.BasePositionRectangle.Height > collisionBox.Y {

			collided = true
			lastPosition = player.BasePositionRectangle.Y

			if collisionBox.Y-player.BasePositionRectangle.Height < collisionPoint {
				collisionPoint = collisionBox.Y - player.BasePositionRectangle.Height
			}
		}
	}

	if collided {
		player.BasePositionRectangle.Y = collisionPoint
		moveCameraDown(player, camera, lastPosition-collisionPoint)
		return nil
	}

	for _, door := range scene.CurrentScene.Doors {
		if player.BasePositionRectangle.X+player.BasePositionRectangle.Width > door.BaseRectangle.X &&
			player.BasePositionRectangle.X < door.BaseRectangle.X+door.BaseRectangle.Width &&
			player.BasePositionRectangle.Y+player.BasePositionRectangle.Height <= door.BaseRectangle.Y &&
			newPlayerRectangleY+player.BasePositionRectangle.Height > door.BaseRectangle.Y {

			return door
		}
	}

	if player.BasePositionRectangle.Y+offset > scene.CurrentScene.Height-global.TileHeight {
		player.BasePositionRectangle.Y = scene.CurrentScene.Height - global.TileHeight
	} else {
		player.BasePositionRectangle.Y += offset
	}

	moveCameraDown(player, camera, offset)
	return nil
}

func moveRight(player *models.Player, camera *rl.Camera2D, offset float32) *models.Door {
	var collided bool
	var collisionPoint float32 = math.MaxFloat32
	var lastPosition float32

	newPlayerRectangleX := player.BasePositionRectangle.X + offset
	for _, collisionBox := range scene.CurrentScene.BaseCollisionBoxes {
		if player.BasePositionRectangle.X+player.BasePositionRectangle.Width <= collisionBox.X &&
			newPlayerRectangleX+player.BasePositionRectangle.Width > collisionBox.X &&
			player.BasePositionRectangle.Y+player.BasePositionRectangle.Height > collisionBox.Y &&
			player.BasePositionRectangle.Y < collisionBox.Y+collisionBox.Height {

			collided = true
			lastPosition = player.BasePositionRectangle.X

			if collisionBox.X-player.BasePositionRectangle.Width < collisionPoint {
				collisionPoint = collisionBox.X - player.BasePositionRectangle.Width
			}
		}
	}

	if collided {
		player.BasePositionRectangle.X = collisionPoint
		moveCameraRight(player, camera, lastPosition-collisionPoint)
		return nil
	}

	for _, door := range scene.CurrentScene.Doors {
		if player.BasePositionRectangle.X+player.BasePositionRectangle.Width <= door.BaseRectangle.X &&
			newPlayerRectangleX+player.BasePositionRectangle.Width > door.BaseRectangle.X &&
			player.BasePositionRectangle.Y+player.BasePositionRectangle.Height > door.BaseRectangle.Y &&
			player.BasePositionRectangle.Y < door.BaseRectangle.Y+door.BaseRectangle.Height {

			return door
		}
	}

	if player.BasePositionRectangle.X+offset > scene.CurrentScene.Width-global.TileWidth {
		player.BasePositionRectangle.X = scene.CurrentScene.Width - global.TileWidth
	} else {
		player.BasePositionRectangle.X += offset
	}

	moveCameraRight(player, camera, offset)
	return nil
}

func moveLeft(player *models.Player, camera *rl.Camera2D, offset float32) *models.Door {
	var collided bool
	var collisionPoint float32
	var lastPosition float32
	newPlayerRectangleX := player.BasePositionRectangle.X - offset
	for _, collisionBox := range scene.CurrentScene.BaseCollisionBoxes {
		if player.BasePositionRectangle.X > collisionBox.X &&
			newPlayerRectangleX <= collisionBox.X+collisionBox.Width &&
			player.BasePositionRectangle.Y+player.BasePositionRectangle.Height > collisionBox.Y &&
			player.BasePositionRectangle.Y < collisionBox.Y+collisionBox.Height {

			collided = true
			lastPosition = player.BasePositionRectangle.X

			if collisionBox.X+collisionBox.Width > collisionPoint {
				collisionPoint = collisionBox.X + collisionBox.Width
			}
		}
	}

	if collided {
		player.BasePositionRectangle.X = collisionPoint
		moveCameraLeft(player, camera, lastPosition-collisionPoint)
		return nil
	}

	for _, door := range scene.CurrentScene.Doors {
		if player.BasePositionRectangle.X > door.BaseRectangle.X &&
			newPlayerRectangleX < door.BaseRectangle.X+door.BaseRectangle.Width &&
			player.BasePositionRectangle.Y+player.BasePositionRectangle.Height > door.BaseRectangle.Y &&
			player.BasePositionRectangle.Y < door.BaseRectangle.Y+door.BaseRectangle.Height {

			return door
		}
	}

	if player.BasePositionRectangle.X-offset < 0 {
		player.BasePositionRectangle.X = 0
	} else {
		player.BasePositionRectangle.X -= offset
	}

	moveCameraLeft(player, camera, offset)
	return nil
}
