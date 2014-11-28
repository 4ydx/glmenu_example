package main

import (
	"fmt"
	glmenu "github.com/4ydx/glmenu"
	gltext "github.com/4ydx/gltext"
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/mathgl/mgl32"
)

func optionMenuInit(window *glfw.Window) (err error) {
	topMargin := float32(50)
	leftMargin := float32(30)
	fontScale := int32(25)
	width, height := window.GetSize()
	optionMenu.Load(float32(width), float32(height), fontScale)
	optionMenu.Font.SetTextLowerBound(0.3)
	optionMenu.ResizeWindow(float32(width), float32(height))
	optionMenu.Background = mgl32.Vec4{1, 1, 1, 1}
	optionMenu.TextScaleRate = 0.05

	var label glmenu.Label
	optionMenu.AddLabel(&label, "Music Volume")

	label.Text.SetColor(0.5, 0.5, 0.5)
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

	var label3 glmenu.Label
	optionMenu.AddLabel(&label3, "Back")

	label3.Text.SetColor(0.5, 0.5, 0.5)
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

	label.Text.SetPosition(-float32(width)/2.0+leftMargin, float32(height)/2.0-topMargin)
	label.Text.Justify(gltext.AlignLeft)
	label3.Text.SetPosition(-float32(width)/2.0+leftMargin, float32(height)/2.0-topMargin-label.Text.Height)
	label3.Text.Justify(gltext.AlignLeft)

	label.AddShadow(1.5, 0, 0, 0)
	label3.AddShadow(1.5, 0, 0, 0)
	return
}
