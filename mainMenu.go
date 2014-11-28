package main

import (
	glmenu "github.com/4ydx/glmenu"
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/mathgl/mgl32"
)

func mainMenuInit(window *glfw.Window) (err error) {
	fontScale := int32(25)
	width, height := window.GetSize()
	mainMenu.Load(float32(width), float32(height), fontScale)
	mainMenu.Font.SetTextLowerBound(0.3)
	mainMenu.ResizeWindow(float32(width), float32(height))
	mainMenu.Background = mgl32.Vec4{0, 0, .20, 0}

	//2DO: sounds

	// start
	var label glmenu.Label
	mainMenu.AddLabel(&label, "Start")

	label.Text.SetColor(0.5, 0.5, 0.5, 1)
	label.OnClick = func(label *glmenu.Label, xPos, yPos float64) (err error) {
		mainMenu.Toggle()
		return
	}
	label.OnHover = func(label *glmenu.Label, xPos, yPos float64) (err error) {
		label.Text.SetColor(0, 250.0/255.0, 154.0/255.0, 1)
		label.Text.AddScale(mainMenu.TextScaleRate)
		return
	}
	label.OnNotHover = func(label *glmenu.Label) (err error) {
		label.Text.SetColor(0.5, 0.5, 0.5, 1)
		label.Text.AddScale(-mainMenu.TextScaleRate)
		return
	}

	// options
	var label2 glmenu.Label
	mainMenu.AddLabel(&label2, "Options")

	label2.Text.SetColor(0.5, 0.5, 0.5, 1)
	label2.OnClick = func(label *glmenu.Label, xPos, yPos float64) (err error) {
		mainMenu.Toggle()
		optionMenu.Toggle()
		return
	}
	label2.OnHover = func(label *glmenu.Label, xPos, yPos float64) (err error) {
		label.Text.SetColor(0, 250.0/255.0, 154.0/255.0, 1)
		label.Text.AddScale(mainMenu.TextScaleRate)
		return
	}
	label2.OnNotHover = func(label *glmenu.Label) (err error) {
		label.Text.SetColor(0.5, 0.5, 0.5, 1)
		label.Text.AddScale(-mainMenu.TextScaleRate)
		return
	}

	// quit
	var label3 glmenu.Label
	mainMenu.AddLabel(&label3, "Quit")

	label3.Text.SetColor(0.5, 0.5, 0.5, 1)
	label3.OnClick = func(label *glmenu.Label, xPos, yPos float64) (err error) {
		window.SetShouldClose(true)
		return
	}
	label3.OnHover = func(label *glmenu.Label, xPos, yPos float64) (err error) {
		label.Text.SetColor(0, 250.0/255.0, 154.0/255.0, 1)
		label.Text.AddScale(mainMenu.TextScaleRate)
		return
	}
	label3.OnNotHover = func(label *glmenu.Label) (err error) {
		label.Text.SetColor(0.5, 0.5, 0.5, 1)
		label.Text.AddScale(-mainMenu.TextScaleRate)
		return
	}

	// simple centering of values
	totalHeight := label.Text.X2.Y - label.Text.X1.Y + label2.Text.X2.Y - label2.Text.X1.Y + label2.Text.X2.Y - label2.Text.X1.Y
	label.Text.SetPosition(0, totalHeight/2)
	label3.Text.SetPosition(0, -totalHeight/2)

	return
}
