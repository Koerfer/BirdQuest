package update

import (
	"BirdQuest/global"
	"BirdQuest/menus"
	"BirdQuest/save"
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
	"time"
)

var lastPressed time.Time

func Menu(player *models.Player, camera rl.Camera2D) (bool, *models.Player, *rl.Camera2D, bool) {
	if rl.IsKeyPressed(rl.KeyEscape) {
		if menus.ActiveMenu != nil && time.Since(lastPressed).Milliseconds() > 200 {
			menus.ActiveMenu.SelectedButton = 0
			menus.ActiveMenu = nil
			return false, nil, nil, false
		}

		lastPressed = time.Now()
		menus.ActiveMenu = menus.AllMenus["main"]
		return true, nil, nil, false
	}

	if menus.ActiveMenu == nil {
		return false, nil, nil, false
	}

	if buttonIndex := checkMouseHover(camera); buttonIndex != nil {
		menus.ActiveMenu.SelectedButton = *buttonIndex
	}

	if rl.IsKeyPressed(rl.KeyDown) || rl.IsKeyPressed(rl.KeyS) {
		menus.ActiveMenu.SelectedButton = (menus.ActiveMenu.SelectedButton + 1) % len(menus.ActiveMenu.Buttons)
	} else if rl.IsKeyPressed(rl.KeyUp) || rl.IsKeyPressed(rl.KeyW) {
		menus.ActiveMenu.SelectedButton = (menus.ActiveMenu.SelectedButton + len(menus.ActiveMenu.Buttons) - 1) % len(menus.ActiveMenu.Buttons)
	} else if rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyEnter) ||
		(checkMouseHover(camera) != nil && rl.IsMouseButtonPressed(rl.MouseButtonLeft)) {
		switch menus.ActiveMenu.Buttons[menus.ActiveMenu.SelectedButton].ActionToDo {
		case menus.ActionInvalid:
			return true, nil, nil, false
		case menus.ActionSave:
			save.Save(player, camera)
			menus.ActiveMenu.Buttons[menus.ActiveMenu.SelectedButton].Name = "DONE"
			go mainMenuSaveButton(menus.ActiveMenu.SelectedButton)
		case menus.ActionLoad:
			p, c := LoadHandler(player, &camera)
			return true, p, c, false
		case menus.ActionExit:
			return false, nil, nil, true
		case menus.ActionResume:
			menus.ActiveMenu = nil
			return true, nil, nil, false
		case menus.ActionOptionsMenu:
			menus.ActiveMenu = menus.AllMenus["options"]
		case menus.ActionMainMenu:
			menus.ActiveMenu = menus.AllMenus["main"]
		case menus.ActionFullScreen:
			if !rl.IsWindowFullscreen() {
				rl.ToggleFullscreen()
				updateDesiredWindowSize(float32(rl.GetMonitorWidth(rl.GetCurrentMonitor())), float32(rl.GetMonitorHeight(rl.GetCurrentMonitor())), player, &camera)
			} else if rl.IsWindowFullscreen() {
				rl.ToggleFullscreen()
				updateDesiredWindowSize(float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight()), player, &camera)
			}
		}
	}

	return true, nil, nil, false
}

func mainMenuSaveButton(button int) {
	time.Sleep(1 * time.Second)
	menus.AllMenus["main"].Buttons[button].Name = "SAVE"
}

func checkMouseHover(camera rl.Camera2D) *int {
	mousePositionAbsolute := rl.GetMousePosition()
	mousePositionRelative := rl.Vector2{
		X: mousePositionAbsolute.X/camera.Zoom + camera.Target.X,
		Y: mousePositionAbsolute.Y/camera.Zoom + camera.Target.Y,
	}

	for i, button := range menus.ActiveMenu.Buttons {
		collisionRectangle := rl.Rectangle{
			X:      button.Rectangle.X/camera.Zoom - button.Rectangle.Width/2/camera.Zoom + camera.Target.X + global.VariableSet.VisibleMapWidth/2,
			Y:      button.Rectangle.Y/camera.Zoom + camera.Target.Y + global.VariableSet.VisibleMapHeight/2 - menus.ActiveMenu.Rectangle.Height/2/camera.Zoom,
			Width:  button.Rectangle.Width / camera.Zoom,
			Height: button.Rectangle.Height / camera.Zoom,
		}
		if collisionRectangle.X < mousePositionRelative.X &&
			collisionRectangle.X+collisionRectangle.Width > mousePositionRelative.X &&
			collisionRectangle.Y < mousePositionRelative.Y &&
			collisionRectangle.Y+collisionRectangle.Height > mousePositionRelative.Y {
			return &i
		}
	}

	return nil
}
