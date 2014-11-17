package main

import (
	glmenu "github.com/4ydx/glmenu"
	glfw "github.com/go-gl/glfw3"
)

func menuInit(window *glfw.Window) (err error) {
	margin := float32(100)
	fontScale := int32(25)
	width, height := window.GetSize()
	menu.Load(float32(width), float32(height), fontScale)
	menu.Font.SetTextLowerBound(0.5)
	menu.ResizeWindow(float32(width), float32(height))

	//2DO: sounds

	// start
	var label glmenu.Label
	menu.AddLabel(&label, "Start")

	label.AddShadow(3, 0, 0, 0)
	label.Text.SetPosition(0, menu.Height-margin)

	label.Text.SetColor(0.5, 0.5, 0.5, 1)
	label.OnClick = func(label *glmenu.Label, xPos, yPos float64) (err error) {
		menu.Toggle()
		return
	}
	label.OnHover = func(label *glmenu.Label, xPos, yPos float64) (err error) {
		label.Text.AddScale(menu.TextScaleRate)
		return
	}
	label.OnNotHover = func(label *glmenu.Label) (err error) {
		label.Text.AddScale(-menu.TextScaleRate)
		return
	}
	label.UpdateShadow(3, 0, 0, 0)

	// options
	var label2 glmenu.Label
	menu.AddLabel(&label2, "Options")

	label2.AddShadow(3, 0, 0, 0)
	label2.Text.SetPosition(0, label.Text.X1.Y)
	label2.Text.SetColor(0.5, 0.5, 0.5, 1)
	label2.OnClick = func(label *glmenu.Label, xPos, yPos float64) (err error) {
		// 2DO: show another menu
		return
	}
	label2.OnHover = func(label *glmenu.Label, xPos, yPos float64) (err error) {
		label.Text.AddScale(menu.TextScaleRate)
		return
	}
	label2.OnNotHover = func(label *glmenu.Label) (err error) {
		label.Text.AddScale(-menu.TextScaleRate)
		return
	}
	label2.UpdateShadow(3, 0, 0, 0)

	// quit
	var label3 glmenu.Label
	menu.AddLabel(&label3, "Quit")

	label3.AddShadow(3, 0, 0, 0)
	label3.Text.SetPosition(0, label2.Text.X1.Y)
	label3.Text.SetColor(0.5, 0.5, 0.5, 1)
	label3.OnClick = func(label *glmenu.Label, xPos, yPos float64) (err error) {
		window.SetShouldClose(true)
		return
	}
	label3.OnHover = func(label *glmenu.Label, xPos, yPos float64) (err error) {
		label.Text.AddScale(menu.TextScaleRate)
		return
	}
	label3.OnNotHover = func(label *glmenu.Label) (err error) {
		label.Text.AddScale(-menu.TextScaleRate)
		return
	}
	label3.UpdateShadow(3, 0, 0, 0)

	return
}
