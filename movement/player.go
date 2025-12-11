package movement

import (
	"BirdQuest/global"
	"BirdQuest/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
	"time"
)

var frameCounter int

func Dash(player *objects.Player, camera *rl.Camera2D, fps int32) {
	player.DashLastUse = time.Now()

	mousePositionAbsolute := rl.GetMousePosition()
	mousePositionRelative := rl.Vector2{
		X: mousePositionAbsolute.X/camera.Zoom + camera.Target.X - (player.Position.X + 16*global.Scale),
		Y: mousePositionAbsolute.Y/camera.Zoom + camera.Target.Y - (player.Position.Y + 16*global.Scale),
	}

	dashDirection := rl.Vector2Normalize(mousePositionRelative)
	if math.Signbit(float64(dashDirection.X)) {
		player.Rotation = float32(360 - math.Acos(float64(-dashDirection.Y))*180/math.Pi)
	} else {
		player.Rotation = float32(math.Acos(float64(-dashDirection.Y)) * 180 / math.Pi)
	}

	speed := 60 / float32(fps) * global.Scale
	dashDirection.X *= 20 * speed
	dashDirection.Y *= 20 * speed
	player.DashDirection = dashDirection
}

func ContinueDash(player *objects.Player, camera *rl.Camera2D, collisionObjects []*rl.Rectangle) {
	player.IsMoving = true
	player.AnimationStep = 1
	player.Rectangle = player.Animation.GetSrc(player.AnimationStep)
	if player.DashDirection.Y < 0 {
		moveUp(player, camera, -player.DashDirection.Y, collisionObjects)
	} else {
		moveDown(player, camera, player.DashDirection.Y, collisionObjects)
	}

	if player.DashDirection.X < 0 {
		moveLeft(player, camera, -player.DashDirection.X, collisionObjects)
	} else {
		moveRight(player, camera, player.DashDirection.X, collisionObjects)
	}

	player.HitBox.X = player.Position.X
	player.HitBox.Y = player.Position.Y
}

func Move(player *objects.Player, camera *rl.Camera2D, fps int32, collisionObjects []*rl.Rectangle) {
	up := rl.IsKeyDown(rl.KeyW)
	down := rl.IsKeyDown(rl.KeyS)
	left := rl.IsKeyDown(rl.KeyA)
	right := rl.IsKeyDown(rl.KeyD)

	speed := 60 / float32(fps) * global.Scale
	diagonalSpeed := speed * 3.535533
	normalSpeed := speed * 5

	if player.Position.Y <= 0 {
		up = false
	}
	if player.Position.Y+32*global.Scale >= global.ScreenHeight*global.Scale {
		down = false
	}
	if player.Position.X+32*global.Scale >= global.ScreenWidth*global.Scale {
		right = false
	}
	if player.Position.X <= 0 {
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

	if up && right {
		player.IsMoving = true
		player.Rotation = 45

		moveUp(player, camera, diagonalSpeed, collisionObjects)
		moveRight(player, camera, diagonalSpeed, collisionObjects)
	} else if up && left {
		player.IsMoving = true
		player.Rotation = 315

		moveUp(player, camera, diagonalSpeed, collisionObjects)
		moveLeft(player, camera, diagonalSpeed, collisionObjects)
	} else if down && left {
		player.IsMoving = true
		player.Rotation = 225

		moveDown(player, camera, diagonalSpeed, collisionObjects)
		moveLeft(player, camera, diagonalSpeed, collisionObjects)
	} else if down && right {
		player.IsMoving = true
		player.Rotation = 135

		moveDown(player, camera, diagonalSpeed, collisionObjects)
		moveRight(player, camera, diagonalSpeed, collisionObjects)
	} else if up {
		player.IsMoving = true
		player.Rotation = 0

		moveUp(player, camera, normalSpeed, collisionObjects)
	} else if down {
		player.IsMoving = true
		player.Rotation = 180
		moveDown(player, camera, normalSpeed, collisionObjects)
	} else if right {
		player.IsMoving = true
		player.Rotation = 90

		moveRight(player, camera, normalSpeed, collisionObjects)
	} else if left {
		player.IsMoving = true
		player.Rotation = 270

		moveLeft(player, camera, normalSpeed, collisionObjects)
	} else {
		player.IsMoving = false
	}

	if player.IsMoving {
		if frameCounter >= int(8/speed) {
			frameCounter = 0
			player.AnimationStep = (player.AnimationStep + 1) % 3
			player.Rectangle = player.Animation.GetSrc(player.AnimationStep)
		}

		frameCounter++
	}

	if !player.IsMoving {
		player.Texture = player.Animation.Texture
		player.Rectangle = player.Animation.GetSrc(7)
		player.AnimationStep = 0
		player.Rotation = 0
		player.IsMoving = false
	}

	player.HitBox.X = player.Position.X
	player.HitBox.Y = player.Position.Y

}

func moveUp(player *objects.Player, camera *rl.Camera2D, offset float32, collisionObjects []*rl.Rectangle) {
	var collided bool
	var collisionPoint float32
	var lastPosition float32

	newPlayerHitBoxY := player.HitBox.Y - offset
	for _, collisionObject := range collisionObjects {
		if player.HitBox.X+player.HitBox.Width > collisionObject.X &&
			player.HitBox.X < collisionObject.X+collisionObject.Width &&
			player.HitBox.Y > collisionObject.Y &&
			newPlayerHitBoxY < collisionObject.Y+collisionObject.Height {

			collided = true
			lastPosition = player.HitBox.Y

			if collisionObject.Y+collisionObject.Height > collisionPoint {
				collisionPoint = collisionObject.Y + collisionObject.Height
			}
		}
	}

	if collided {
		player.Position.Y = collisionPoint
		moveCameraUp(player, camera, lastPosition-collisionPoint)
		return
	}

	if player.Position.Y-offset < 0 {
		player.Position.Y = 0
	} else {
		player.Position.Y -= offset
	}

	moveCameraUp(player, camera, offset)
}

func moveDown(player *objects.Player, camera *rl.Camera2D, offset float32, collisionObjects []*rl.Rectangle) {
	var collided bool
	var collisionPoint float32 = math.MaxFloat32
	var lastPosition float32

	newPlayerHitBoxY := player.HitBox.Y + offset
	for _, collisionObject := range collisionObjects {
		if player.HitBox.X+player.HitBox.Width > collisionObject.X &&
			player.HitBox.X < collisionObject.X+collisionObject.Width &&
			player.HitBox.Y+player.HitBox.Height <= collisionObject.Y &&
			newPlayerHitBoxY+player.HitBox.Height > collisionObject.Y {

			collided = true
			lastPosition = player.HitBox.Y

			if collisionObject.Y-player.HitBox.Height < collisionPoint {
				collisionPoint = collisionObject.Y - player.HitBox.Height
			}
		}
	}

	if collided {
		player.Position.Y = collisionPoint
		moveCameraDown(player, camera, lastPosition-collisionPoint)
		return
	}

	if player.Position.Y+offset > (global.ScreenHeight-32)*global.Scale {
		player.Position.Y = (global.ScreenHeight - 32) * global.Scale
	} else {
		player.Position.Y += offset
	}

	moveCameraDown(player, camera, offset)
}

func moveRight(player *objects.Player, camera *rl.Camera2D, offset float32, collisionObjects []*rl.Rectangle) {
	var collided bool
	var collisionPoint float32 = math.MaxFloat32
	var lastPosition float32

	newPlayerHitBoxX := player.HitBox.X + offset
	for _, collisionObject := range collisionObjects {
		if player.HitBox.X+player.HitBox.Width <= collisionObject.X &&
			newPlayerHitBoxX+player.HitBox.Width > collisionObject.X &&
			player.HitBox.Y+player.HitBox.Height > collisionObject.Y &&
			player.HitBox.Y < collisionObject.Y+collisionObject.Height {

			collided = true
			lastPosition = player.HitBox.X

			if collisionObject.X-player.HitBox.Width < collisionPoint {
				collisionPoint = collisionObject.X - player.HitBox.Width
			}
		}
	}

	if collided {
		player.Position.X = collisionPoint
		moveCameraRight(player, camera, lastPosition-collisionPoint)
		return
	}

	if player.Position.X+offset > (global.ScreenWidth-32)*global.Scale {
		player.Position.X = (global.ScreenWidth - 32) * global.Scale
	} else {
		player.Position.X += offset
	}

	moveCameraRight(player, camera, offset)
}

func moveLeft(player *objects.Player, camera *rl.Camera2D, offset float32, collisionObjects []*rl.Rectangle) {
	var collided bool
	var collisionPoint float32
	var lastPosition float32
	newPlayerHitBoxX := player.HitBox.X - offset
	for _, collisionObject := range collisionObjects {
		if player.HitBox.X > collisionObject.X &&
			newPlayerHitBoxX < collisionObject.X+collisionObject.Width &&
			player.HitBox.Y+player.HitBox.Height > collisionObject.Y &&
			player.HitBox.Y < collisionObject.Y+collisionObject.Height {

			collided = true
			lastPosition = player.HitBox.X

			if collisionObject.X+collisionObject.Width > collisionPoint {
				collisionPoint = collisionObject.X + collisionObject.Width
			}
		}
	}

	if collided {
		player.Position.X = collisionPoint
		moveCameraLeft(player, camera, lastPosition-collisionPoint)
		return
	}

	if player.Position.X-offset < 0 {
		player.Position.X = 0
	} else {
		player.Position.X -= offset
	}

	moveCameraLeft(player, camera, offset)
}
