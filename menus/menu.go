package menus

import (
	"BirdQuest/global"
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

type Menu struct {
	Buttons       []*Button
	ButtonHeight  float32
	ButtonSpacing float32
	BaseFontSize  float32
	FontSize      float32

	BaseRectangle  *rl.Rectangle
	Rectangle      *rl.Rectangle
	SelectedButton int
}

type Button struct {
	Name          string
	BaseRectangle *rl.Rectangle
	Rectangle     *rl.Rectangle

	DefaultColour  rl.Color
	SelectedColour rl.Color
	ActionToDo     Action
}

type Action int8

const (
	ActionInvalid Action = iota
	ActionSave
	ActionLoad
	ActionExit
	ActionResume
)

var ActiveMenu *Menu
var AllMenus map[string]*Menu

func PrepareMenus() {
	AllMenus = make(map[string]*Menu)

	var standardWidth float32 = 280
	var standardButtonHeight float32 = 95
	var standardButtonSpacing float32 = 8

	mainMenu := CreateMenu(standardWidth, standardButtonHeight, standardButtonSpacing, 64)
	mainMenu.AddButton("START", rl.Green, rl.White, ActionResume)
	mainMenu.AddButton("OPTIONS", rl.Green, rl.White, ActionInvalid)
	mainMenu.AddButton("SAVE", rl.Green, rl.White, ActionSave)
	mainMenu.AddButton("LOAD", rl.Green, rl.White, ActionLoad)
	mainMenu.AddButton("EXIT", rl.Green, rl.White, ActionExit)

	AllMenus["main"] = mainMenu
}

func CreateMenu(menuWidth, buttonHeight, buttonSpacing, fontSize float32) *Menu {
	menu := &Menu{
		Buttons:       make([]*Button, 0),
		ButtonHeight:  buttonHeight,
		ButtonSpacing: buttonSpacing,
		BaseFontSize:  fontSize,
		FontSize:      fontSize * global.VariableSet.EntityScale,
		BaseRectangle: &rl.Rectangle{
			X:      0,
			Y:      0,
			Width:  menuWidth,
			Height: buttonSpacing,
		},
		Rectangle: &rl.Rectangle{
			X:      0,
			Y:      0,
			Width:  menuWidth * global.VariableSet.EntityScale,
			Height: buttonSpacing * global.VariableSet.EntityScale,
		},
		SelectedButton: 0,
	}

	return menu
}

func (menu *Menu) AddButton(name string, selectedColour, defaultColour rl.Color, action Action) {
	menu.BaseRectangle.Height += menu.ButtonHeight + menu.ButtonSpacing
	menu.Rectangle.Height = menu.BaseRectangle.Height * global.VariableSet.EntityScale

	button := &Button{}
	var buttonSpacingFromMenu float32 = 10
	button.BaseRectangle = &rl.Rectangle{
		X:      menu.Rectangle.X + buttonSpacingFromMenu - menu.ButtonSpacing,
		Y:      float32(len(menu.Buttons))*(menu.ButtonHeight) + float32(len(menu.Buttons)+1)*menu.ButtonSpacing,
		Width:  menu.BaseRectangle.Width - buttonSpacingFromMenu*2,
		Height: menu.ButtonHeight,
	}
	button.Rectangle = &rl.Rectangle{
		X:      button.BaseRectangle.X * global.VariableSet.EntityScale,
		Y:      button.BaseRectangle.Y * global.VariableSet.EntityScale,
		Width:  button.BaseRectangle.Width * global.VariableSet.EntityScale,
		Height: button.BaseRectangle.Height * global.VariableSet.EntityScale,
	}

	button.Name = name
	button.SelectedColour = selectedColour
	button.DefaultColour = defaultColour
	button.ActionToDo = action

	menu.Buttons = append(menu.Buttons, button)
}

func (menu *Menu) Draw(camera rl.Camera2D) {
	rl.DrawRectanglePro(rl.Rectangle{
		X:      menu.Rectangle.X / camera.Zoom,
		Y:      menu.Rectangle.Y / camera.Zoom,
		Width:  menu.Rectangle.Width / camera.Zoom,
		Height: menu.Rectangle.Height / camera.Zoom,
	},
		rl.Vector2{
			X: -camera.Target.X - global.VariableSet.VisibleMapWidth/2 + menu.Rectangle.Width/2/camera.Zoom,
			Y: -camera.Target.Y - global.VariableSet.VisibleMapHeight/2 + menu.Rectangle.Height/2/camera.Zoom,
		},
		0,
		color.RGBA{A: 200})

	for i, button := range menu.Buttons {
		colour := button.DefaultColour
		if i == menu.SelectedButton {
			colour = button.SelectedColour
		}

		newRectangle := rl.Rectangle{
			X:      button.Rectangle.X/camera.Zoom - button.Rectangle.Width/2/camera.Zoom + camera.Target.X + global.VariableSet.VisibleMapWidth/2,
			Y:      button.Rectangle.Y/camera.Zoom + camera.Target.Y + global.VariableSet.VisibleMapHeight/2 - menu.Rectangle.Height/2/camera.Zoom,
			Width:  button.Rectangle.Width / camera.Zoom,
			Height: button.Rectangle.Height / camera.Zoom,
		}
		rl.DrawRectangleLinesEx(newRectangle, 2*global.VariableSet.EntityScale/camera.Zoom, colour)

		textSize := rl.MeasureTextEx(global.Font,
			button.Name, menu.FontSize/camera.Zoom, 2*global.VariableSet.EntityScale/camera.Zoom)

		rl.DrawTextPro(global.Font,
			button.Name,
			rl.NewVector2(
				newRectangle.X+button.Rectangle.Width/2/camera.Zoom,
				newRectangle.Y+button.Rectangle.Height/2/camera.Zoom,
			),
			rl.NewVector2(textSize.X/2, textSize.Y/2),
			0, menu.FontSize/camera.Zoom, 2*global.VariableSet.EntityScale/camera.Zoom, colour)
	}
}
