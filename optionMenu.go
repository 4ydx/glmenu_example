package main

import (
	"fmt"
	glmenu "github.com/4ydx/glmenu"
	glfw "github.com/go-gl/glfw3"
)

// 2DO: figure out how to right align the text... it
//      is automatically centered right now...
func optionMenuInit(window *glfw.Window) (err error) {
	margin := float32(100)
	fontScale := int32(25)
	width, height := window.GetSize()
	optionMenu.Load(float32(width), float32(height), fontScale)
	optionMenu.Font.SetTextLowerBound(0.3)
	optionMenu.ResizeWindow(float32(width), float32(height))

	var label glmenu.Label
	optionMenu.AddLabel(&label, "Music Volume")

	label.Text.SetColor(0.5, 0.5, 0.5, 1)
	label.Text.SetPosition(0, 0)
	label.OnClick = func(label *glmenu.Label, xPos, yPos float64) (err error) {
		fmt.Println("clicked", xPos, yPos)
		return
	}
	label.OnHover = func(label *glmenu.Label, xPos, yPos float64) (err error) {
		label.Text.AddScale(optionMenu.TextScaleRate)
		return
	}
	label.OnNotHover = func(label *glmenu.Label) (err error) {
		label.Text.AddScale(-optionMenu.TextScaleRate)
		return
	}
	label.AddShadow(1.5, 0, 0, 0)

	var label3 glmenu.Label
	optionMenu.AddLabel(&label3, "Back")

	label3.AddShadow(1.5, 0, 0, 0)
	label3.Text.SetColor(0.5, 0.5, 0.5, 1)
	label3.Text.SetPosition(0, label.Text.X1.Y-margin)
	label3.OnClick = func(label *glmenu.Label, xPos, yPos float64) (err error) {
		optionMenu.Toggle()
		mainMenu.Toggle()
		return
	}
	label3.OnHover = func(label *glmenu.Label, xPos, yPos float64) (err error) {
		label.Text.AddScale(optionMenu.TextScaleRate)
		return
	}
	label3.OnNotHover = func(label *glmenu.Label) (err error) {
		label.Text.AddScale(-optionMenu.TextScaleRate)
		return
	}
	label3.AddShadow(1.5, 0, 0, 0)

	return
}
