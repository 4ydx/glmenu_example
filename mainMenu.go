package main

import (
	"fmt"
	glmenu "github.com/4ydx/glmenu"
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/mathgl/mgl32"
	"os"
)

func mainMenuInit(window *glfw.Window) (err error) {
	fontScale := int32(25)
	width, height := window.GetSize()
	err = mainMenu.Load(float32(width), float32(height), fontScale)
	if err != nil {
		fmt.Println("error loading the font")
		os.Exit(1)
	}
	mainMenu.ResizeWindow(float32(width), float32(height))
	mainMenu.Background = mgl32.Vec4{0, 0, .20, 0}

	//2DO: sounds

	// start
	var textbox1 glmenu.TextBox
	mainMenu.AddTextBox(&textbox1, "127.0.0.1")
	textbox1.Text.SetColor(1, 1, 1)

	// options
	var label2 glmenu.Label
	mainMenu.AddLabel(&label2, "Options")

	label2.Text.SetColor(0.5, 0.5, 0.5)
	label2.OnClick = func(label *glmenu.Label, xPos, yPos float64, button glmenu.MouseClick, inBox bool) {
		label.Text.SetColor(250.0/255.0, 0, 154.0/255.0)
	}
	label2.OnRelease = func(label *glmenu.Label, xPos, yPos float64, button glmenu.MouseClick, inBox bool) {
		label.Text.SetColor(0, 250.0/255.0, 154.0/255.0)
		if inBox {
			mainMenu.Toggle()
			optionMenu.Toggle()
		}
		if label.IsHover {
			label.Text.SetColor(0, 250.0/255.0, 154.0/255.0)
		} else {
			label.Text.SetColor(0.5, 0.5, 0.5)
		}
	}
	label2.OnHover = func(label *glmenu.Label, xPos, yPos float64, button glmenu.MouseClick, inBox bool) {
		if !label.IsClick {
			label.Text.SetColor(0, 250.0/255.0, 154.0/255.0)
			label.Text.AddScale(mainMenu.TextScaleRate)
		}
	}
	label2.OnNotHover = func(label *glmenu.Label) {
		if !label.IsClick {
			label.Text.SetColor(0.5, 0.5, 0.5)
			label.Text.AddScale(-mainMenu.TextScaleRate)
		}
	}

	// quit
	var label3 glmenu.Label
	mainMenu.AddLabel(&label3, "Quit")

	label3.Text.SetColor(0.5, 0.5, 0.5)
	label3.OnClick = func(label *glmenu.Label, xPos, yPos float64, button glmenu.MouseClick, inBox bool) {
		label.Text.SetColor(250.0/255.0, 0, 154.0/255.0)
	}
	label3.OnRelease = func(label *glmenu.Label, xPos, yPos float64, button glmenu.MouseClick, inBox bool) {
		label.Text.SetColor(0, 250.0/255.0, 154.0/255.0)
		if inBox {
			window.SetShouldClose(true)
		}
		if label.IsHover {
			label.Text.SetColor(0, 250.0/255.0, 154.0/255.0)
		} else {
			label.Text.SetColor(0.5, 0.5, 0.5)
		}
	}
	label3.OnHover = func(label *glmenu.Label, xPos, yPos float64, button glmenu.MouseClick, inBox bool) {
		if !label.IsClick {
			label.Text.SetColor(0, 250.0/255.0, 154.0/255.0)
			label.Text.AddScale(mainMenu.TextScaleRate)
		}
	}
	label3.OnNotHover = func(label *glmenu.Label) {
		if !label.IsClick {
			label.Text.SetColor(0.5, 0.5, 0.5)
			label.Text.AddScale(-mainMenu.TextScaleRate)
		}
	}

	// simple centering of values
	totalHeight := textbox1.Text.X2.Y - textbox1.Text.X1.Y +
		label2.Text.X2.Y - label2.Text.X1.Y +
		label2.Text.X2.Y - label2.Text.X1.Y
	textbox1.Text.SetPosition(0, totalHeight/2)

	label3.Text.SetPosition(0, -totalHeight/2)

	return
}
