package main

import (
	"fmt"
	"github.com/4ydx/glmenu"
	"github.com/4ydx/gltext"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"os"
)

func optionMenuInit(window *glfw.Window) (err error) {
	topMargin := float32(50)
	leftMargin := float32(30)
	fontScale := int32(25)
	width, height := window.GetSize()
	err = optionMenu.Load(float32(width), float32(height), fontScale, mgl32.Vec2{})
	if err != nil {
		fmt.Println("error loading font")
		os.Exit(1)
	}
	optionMenu.ResizeWindow(float32(width), float32(height))
	optionMenu.Background = mgl32.Vec4{1, 1, 1, 1}
	optionMenu.TextScaleRate = 0.05

	var label1 glmenu.Label
	optionMenu.AddLabel(&label1, "Music Volume")

	label1.Text.SetColor(0.5, 0.5, 0.5)
	label1.OnClick = func(label *glmenu.Label, xPos, yPos float64, button glmenu.MouseClick, isBox bool) {
		fmt.Println("clicked", xPos, yPos)
	}
	label1.OnHover = func(label *glmenu.Label, xPos, yPos float64, button glmenu.MouseClick, isBox bool) {
		label.Text.AddScale(optionMenu.TextScaleRate)
	}
	label1.OnNotHover = func(label *glmenu.Label) {
		label.Text.AddScale(-optionMenu.TextScaleRate)
	}

	var label3 glmenu.Label
	optionMenu.AddLabel(&label3, "Back")

	label3.Text.SetColor(0.5, 0.5, 0.5)
	label3.OnClick = func(label *glmenu.Label, xPos, yPos float64, button glmenu.MouseClick, isBox bool) {
		optionMenu.Toggle()
		mainMenu.Toggle()
	}
	label3.OnHover = func(label *glmenu.Label, xPos, yPos float64, button glmenu.MouseClick, isBox bool) {
		label.Text.AddScale(optionMenu.TextScaleRate)
	}
	label3.OnNotHover = func(label *glmenu.Label) {
		label.Text.AddScale(-optionMenu.TextScaleRate)
	}

	label1.Text.SetPosition(-float32(width)/2.0+leftMargin, float32(height)/2.0-topMargin)
	label1.Text.Justify(gltext.AlignLeft)
	label3.Text.SetPosition(-float32(width)/2.0+leftMargin, float32(height)/2.0-topMargin-label1.Text.Height)
	label3.Text.Justify(gltext.AlignLeft)

	label1.AddShadow(1.5, 0, 0, 0)
	label3.AddShadow(1.5, 0, 0, 0)
	return
}
