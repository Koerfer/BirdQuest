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
		X: mousePositionAbsolute.X/camera.Zoom + camera.Target.X - (player.Rectangle.X + global.VariableSet.PlayerMiddleOffset),
		Y: mousePositionAbsolute.Y/camera.Zoom + camera.Target.Y - (player.Rectangle.Y + global.VariableSet.PlayerMiddleOffset),
	}

	dashDirection := rl.Vector2Normalize(mousePositionRelative)
	if math.Signbit(float64(dashDirection.X)) {
		player.Rotation = float32(360 - math.Acos(float64(-dashDirection.Y))*180/math.Pi)
	} else {
		player.Rotation = float32(math.Acos(float64(-dashDirection.Y)) * 180 / math.Pi)
	}

	dashDirection.X *= 20 * global.VariableSet.Speed
	dashDirection.Y *= 20 * global.VariableSet.Speed
	player.DashDirection = dashDirection
}

func ContinueDash(player *models.Player, camera *rl.Camera2D) *models.Door {
	player.IsMoving = true
	player.AnimationStep = 1
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

	diagonalSpeed := global.VariableSet.Speed * 3.535533
	normalSpeed := global.VariableSet.Speed * 5

	if player.Rectangle.Y <= 0 {
		up = false
	}
	if player.Rectangle.Y+global.VariableSet.EntitySize >= global.VariableSet.MapHeight {
		down = false
	}
	if player.Rectangle.X+global.VariableSet.EntitySize >= global.VariableSet.MapWidth {
		right = false
	}
	if player.Rectangle.X <= 0 {
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
		if frameCounter >= int(8/global.VariableSet.FpsScale) {
			frameCounter = 0
			player.AnimationStep = (player.AnimationStep + 1) % 3
			player.BaseRectangle = player.Animation.GetRectangleAreaInTexture(player.AnimationStep)
		}

		frameCounter++
	}

	if !player.IsMoving {
		player.Texture = player.Animation.Texture
		player.BaseRectangle = player.Animation.GetRectangleAreaInTexture(7)
		player.AnimationStep = 0
		player.Rotation = 0
		player.IsMoving = false
	}

	player.Rectangle.X = player.Rectangle.X
	player.Rectangle.Y = player.Rectangle.Y

	return goThroughDoor
}

func moveUp(player *models.Player, camera *rl.Camera2D, offset float32) *models.Door {
	var collided bool
	var collisionPoint float32
	var lastPosition float32

	newPlayerRectangleY := player.Rectangle.Y - offset
	for _, collisionBox := range scene.CurrentScene.CollisionBoxes {
		if player.Rectangle.X+player.Rectangle.Width > collisionBox.X &&
			player.Rectangle.X < collisionBox.X+collisionBox.Width &&
			player.Rectangle.Y > collisionBox.Y &&
			newPlayerRectangleY < collisionBox.Y+collisionBox.Height {

			collided = true
			lastPosition = player.Rectangle.Y

			if collisionBox.Y+collisionBox.Height > collisionPoint {
				collisionPoint = collisionBox.Y + collisionBox.Height
			}
		}
	}

	if collided {
		player.Rectangle.Y = collisionPoint
		moveCameraUp(player, camera, lastPosition-collisionPoint)
		return nil
	}

	for _, door := range scene.CurrentScene.Doors {
		if player.Rectangle.X+player.Rectangle.Width > door.Rectangle.X &&
			player.Rectangle.X < door.Rectangle.X+door.Rectangle.Width &&
			player.Rectangle.Y > door.Rectangle.Y &&
			newPlayerRectangleY < door.Rectangle.Y+door.Rectangle.Height {

			return door
		}
	}

	if player.Rectangle.Y-offset < 0 {
		player.Rectangle.Y = 0
	} else {
		player.Rectangle.Y -= offset
	}

	moveCameraUp(player, camera, offset)
	return nil
}

func moveDown(player *models.Player, camera *rl.Camera2D, offset float32) *models.Door {
	var collided bool
	var collisionPoint float32 = math.MaxFloat32
	var lastPosition float32

	newPlayerRectangleY := player.Rectangle.Y + offset
	for _, collisionBox := range scene.CurrentScene.CollisionBoxes {
		if player.Rectangle.X+player.Rectangle.Width > collisionBox.X &&
			player.Rectangle.X < collisionBox.X+collisionBox.Width &&
			player.Rectangle.Y+player.Rectangle.Height <= collisionBox.Y &&
			newPlayerRectangleY+player.Rectangle.Height > collisionBox.Y {

			collided = true
			lastPosition = player.Rectangle.Y

			if collisionBox.Y-player.Rectangle.Height < collisionPoint {
				collisionPoint = collisionBox.Y - player.Rectangle.Height
			}
		}
	}

	if collided {
		player.Rectangle.Y = collisionPoint
		moveCameraDown(player, camera, lastPosition-collisionPoint)
		return nil
	}

	for _, door := range scene.CurrentScene.Doors {
		if player.Rectangle.X+player.Rectangle.Width > door.Rectangle.X &&
			player.Rectangle.X < door.Rectangle.X+door.Rectangle.Width &&
			player.Rectangle.Y+player.Rectangle.Height <= door.Rectangle.Y &&
			newPlayerRectangleY+player.Rectangle.Height > door.Rectangle.Y {

			return door
		}
	}

	if player.Rectangle.Y+offset > global.VariableSet.MapHeight-global.VariableSet.EntitySize {
		player.Rectangle.Y = global.VariableSet.MapHeight - global.VariableSet.EntitySize
	} else {
		player.Rectangle.Y += offset
	}

	moveCameraDown(player, camera, offset)
	return nil
}

func moveRight(player *models.Player, camera *rl.Camera2D, offset float32) *models.Door {
	var collided bool
	var collisionPoint float32 = math.MaxFloat32
	var lastPosition float32

	newPlayerRectangleX := player.Rectangle.X + offset
	for _, collisionBox := range scene.CurrentScene.CollisionBoxes {
		if player.Rectangle.X+player.Rectangle.Width <= collisionBox.X &&
			newPlayerRectangleX+player.Rectangle.Width > collisionBox.X &&
			player.Rectangle.Y+player.Rectangle.Height > collisionBox.Y &&
			player.Rectangle.Y < collisionBox.Y+collisionBox.Height {

			collided = true
			lastPosition = player.Rectangle.X

			if collisionBox.X-player.Rectangle.Width < collisionPoint {
				collisionPoint = collisionBox.X - player.Rectangle.Width
			}
		}
	}

	if collided {
		player.Rectangle.X = collisionPoint
		moveCameraRight(player, camera, lastPosition-collisionPoint)
		return nil
	}

	for _, door := range scene.CurrentScene.Doors {
		if player.Rectangle.X+player.Rectangle.Width <= door.Rectangle.X &&
			newPlayerRectangleX+player.Rectangle.Width > door.Rectangle.X &&
			player.Rectangle.Y+player.Rectangle.Height > door.Rectangle.Y &&
			player.Rectangle.Y < door.Rectangle.Y+door.Rectangle.Height {

			return door
		}
	}

	if player.Rectangle.X+offset > global.VariableSet.MapWidth-global.VariableSet.EntitySize {
		player.Rectangle.X = global.VariableSet.MapWidth - global.VariableSet.EntitySize
	} else {
		player.Rectangle.X += offset
	}

	moveCameraRight(player, camera, offset)
	return nil
}

func moveLeft(player *models.Player, camera *rl.Camera2D, offset float32) *models.Door {
	var collided bool
	var collisionPoint float32
	var lastPosition float32
	newPlayerRectangleX := player.Rectangle.X - offset
	for _, collisionBox := range scene.CurrentScene.CollisionBoxes {
		if player.Rectangle.X > collisionBox.X &&
			newPlayerRectangleX < collisionBox.X+collisionBox.Width &&
			player.Rectangle.Y+player.Rectangle.Height > collisionBox.Y &&
			player.Rectangle.Y < collisionBox.Y+collisionBox.Height {

			collided = true
			lastPosition = player.Rectangle.X

			if collisionBox.X+collisionBox.Width > collisionPoint {
				collisionPoint = collisionBox.X + collisionBox.Width
			}
		}
	}

	if collided {
		player.Rectangle.X = collisionPoint
		moveCameraLeft(player, camera, lastPosition-collisionPoint)
		return nil
	}

	for _, door := range scene.CurrentScene.Doors {
		if player.Rectangle.X > door.Rectangle.X &&
			newPlayerRectangleX < door.Rectangle.X+door.Rectangle.Width &&
			player.Rectangle.Y+player.Rectangle.Height > door.Rectangle.Y &&
			player.Rectangle.Y < door.Rectangle.Y+door.Rectangle.Height {

			return door
		}
	}

	if player.Rectangle.X-offset < 0 {
		player.Rectangle.X = 0
	} else {
		player.Rectangle.X -= offset
	}

	moveCameraLeft(player, camera, offset)
	return nil
}
