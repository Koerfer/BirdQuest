package movement

import (
	"BirdQuest/global"
	"BirdQuest/scene"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
	"time"
)

var frameCounter int

func Dash(player *scene.Player, camera *rl.Camera2D) {
	player.DashLastUse = time.Now()

	mousePositionAbsolute := rl.GetMousePosition()
	mousePositionRelative := rl.Vector2{
		X: mousePositionAbsolute.X/camera.Zoom + camera.Target.X - (player.Position.X + global.VariableSet.PlayerMiddleOffset),
		Y: mousePositionAbsolute.Y/camera.Zoom + camera.Target.Y - (player.Position.Y + global.VariableSet.PlayerMiddleOffset),
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

func ContinueDash(player *scene.Player, camera *rl.Camera2D) {
	player.IsMoving = true
	player.AnimationStep = 1
	player.Rectangle = player.Animation.GetRectangleAreaInTexture(player.AnimationStep)
	if player.DashDirection.Y < 0 {
		moveUp(player, camera, -player.DashDirection.Y)
	} else {
		moveDown(player, camera, player.DashDirection.Y)
	}

	if player.DashDirection.X < 0 {
		moveLeft(player, camera, -player.DashDirection.X)
	} else {
		moveRight(player, camera, player.DashDirection.X)
	}

	player.HitBox.X = player.Position.X
	player.HitBox.Y = player.Position.Y
}

func Move(player *scene.Player, camera *rl.Camera2D) {
	up := rl.IsKeyDown(rl.KeyW)
	down := rl.IsKeyDown(rl.KeyS)
	left := rl.IsKeyDown(rl.KeyA)
	right := rl.IsKeyDown(rl.KeyD)

	diagonalSpeed := global.VariableSet.Speed * 3.535533
	normalSpeed := global.VariableSet.Speed * 5

	if player.Position.Y <= 0 {
		up = false
	}
	if player.Position.Y+global.VariableSet.EntitySize >= global.VariableSet.MapHeight {
		down = false
	}
	if player.Position.X+global.VariableSet.EntitySize >= global.VariableSet.MapWidth {
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

		moveUp(player, camera, diagonalSpeed)
		moveRight(player, camera, diagonalSpeed)
	} else if up && left {
		player.IsMoving = true
		player.Rotation = 315

		moveUp(player, camera, diagonalSpeed)
		moveLeft(player, camera, diagonalSpeed)
	} else if down && left {
		player.IsMoving = true
		player.Rotation = 225

		moveDown(player, camera, diagonalSpeed)
		moveLeft(player, camera, diagonalSpeed)
	} else if down && right {
		player.IsMoving = true
		player.Rotation = 135

		moveDown(player, camera, diagonalSpeed)
		moveRight(player, camera, diagonalSpeed)
	} else if up {
		player.IsMoving = true
		player.Rotation = 0

		moveUp(player, camera, normalSpeed)
	} else if down {
		player.IsMoving = true
		player.Rotation = 180
		moveDown(player, camera, normalSpeed)
	} else if right {
		player.IsMoving = true
		player.Rotation = 90

		moveRight(player, camera, normalSpeed)
	} else if left {
		player.IsMoving = true
		player.Rotation = 270

		moveLeft(player, camera, normalSpeed)
	} else {
		player.IsMoving = false
	}

	if player.IsMoving {
		if frameCounter >= int(8/global.VariableSet.FpsScale) {
			frameCounter = 0
			player.AnimationStep = (player.AnimationStep + 1) % 3
			player.Rectangle = player.Animation.GetRectangleAreaInTexture(player.AnimationStep)
		}

		frameCounter++
	}

	if !player.IsMoving {
		player.Texture = player.Animation.Texture
		player.Rectangle = player.Animation.GetRectangleAreaInTexture(7)
		player.AnimationStep = 0
		player.Rotation = 0
		player.IsMoving = false
	}

	player.HitBox.X = player.Position.X
	player.HitBox.Y = player.Position.Y

}

func moveUp(player *scene.Player, camera *rl.Camera2D, offset float32) {
	var collided bool
	var collisionPoint float32
	var lastPosition float32

	newPlayerHitBoxY := player.HitBox.Y - offset
	for _, collisionBox := range scene.CurrentScene.CollisionBoxes {
		if player.HitBox.X+player.HitBox.Width > collisionBox.X &&
			player.HitBox.X < collisionBox.X+collisionBox.Width &&
			player.HitBox.Y > collisionBox.Y &&
			newPlayerHitBoxY < collisionBox.Y+collisionBox.Height {

			collided = true
			lastPosition = player.HitBox.Y

			if collisionBox.Y+collisionBox.Height > collisionPoint {
				collisionPoint = collisionBox.Y + collisionBox.Height
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

func moveDown(player *scene.Player, camera *rl.Camera2D, offset float32) {
	var collided bool
	var collisionPoint float32 = math.MaxFloat32
	var lastPosition float32

	newPlayerHitBoxY := player.HitBox.Y + offset
	for _, collisionBox := range scene.CurrentScene.CollisionBoxes {
		if player.HitBox.X+player.HitBox.Width > collisionBox.X &&
			player.HitBox.X < collisionBox.X+collisionBox.Width &&
			player.HitBox.Y+player.HitBox.Height <= collisionBox.Y &&
			newPlayerHitBoxY+player.HitBox.Height > collisionBox.Y {

			collided = true
			lastPosition = player.HitBox.Y

			if collisionBox.Y-player.HitBox.Height < collisionPoint {
				collisionPoint = collisionBox.Y - player.HitBox.Height
			}
		}
	}

	if collided {
		player.Position.Y = collisionPoint
		moveCameraDown(player, camera, lastPosition-collisionPoint)
		return
	}

	if player.Position.Y+offset > global.VariableSet.MapHeight-global.VariableSet.EntitySize {
		player.Position.Y = global.VariableSet.MapHeight - global.VariableSet.EntitySize
	} else {
		player.Position.Y += offset
	}

	moveCameraDown(player, camera, offset)
}

func moveRight(player *scene.Player, camera *rl.Camera2D, offset float32) {
	var collided bool
	var collisionPoint float32 = math.MaxFloat32
	var lastPosition float32

	newPlayerHitBoxX := player.HitBox.X + offset
	for _, collisionBox := range scene.CurrentScene.CollisionBoxes {
		if player.HitBox.X+player.HitBox.Width <= collisionBox.X &&
			newPlayerHitBoxX+player.HitBox.Width > collisionBox.X &&
			player.HitBox.Y+player.HitBox.Height > collisionBox.Y &&
			player.HitBox.Y < collisionBox.Y+collisionBox.Height {

			collided = true
			lastPosition = player.HitBox.X

			if collisionBox.X-player.HitBox.Width < collisionPoint {
				collisionPoint = collisionBox.X - player.HitBox.Width
			}
		}
	}

	if collided {
		player.Position.X = collisionPoint
		moveCameraRight(player, camera, lastPosition-collisionPoint)
		return
	}

	if player.Position.X+offset > global.VariableSet.MapWidth-global.VariableSet.EntitySize {
		player.Position.X = global.VariableSet.MapWidth - global.VariableSet.EntitySize
	} else {
		player.Position.X += offset
	}

	moveCameraRight(player, camera, offset)
}

func moveLeft(player *scene.Player, camera *rl.Camera2D, offset float32) {
	var collided bool
	var collisionPoint float32
	var lastPosition float32
	newPlayerHitBoxX := player.HitBox.X - offset
	for _, collisionBox := range scene.CurrentScene.CollisionBoxes {
		if player.HitBox.X > collisionBox.X &&
			newPlayerHitBoxX < collisionBox.X+collisionBox.Width &&
			player.HitBox.Y+player.HitBox.Height > collisionBox.Y &&
			player.HitBox.Y < collisionBox.Y+collisionBox.Height {

			collided = true
			lastPosition = player.HitBox.X

			if collisionBox.X+collisionBox.Width > collisionPoint {
				collisionPoint = collisionBox.X + collisionBox.Width
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
