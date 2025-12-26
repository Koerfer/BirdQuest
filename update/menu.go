package update

import (
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

	if rl.IsKeyPressed(rl.KeyDown) || rl.IsKeyPressed(rl.KeyS) {
		menus.ActiveMenu.SelectedButton = (menus.ActiveMenu.SelectedButton + 1) % len(menus.ActiveMenu.Buttons)
	} else if rl.IsKeyPressed(rl.KeyUp) || rl.IsKeyPressed(rl.KeyW) {
		menus.ActiveMenu.SelectedButton = (menus.ActiveMenu.SelectedButton + len(menus.ActiveMenu.Buttons) - 1) % len(menus.ActiveMenu.Buttons)
	} else if rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyEnter) {
		switch menus.ActiveMenu.Buttons[menus.ActiveMenu.SelectedButton].ActionToDo {
		case menus.ActionInvalid:
			//do nothing
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
