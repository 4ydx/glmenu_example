package main

import (
	glmenu "github.com/4ydx/glmenu"
)

func menuInit() (err error) {
	// menu settings
	width, height := window.GetSize()

	margin := float32(60)
	fontScale := int32(25)
	menu.Load(400, 350, fontScale)
	menu.Font.SetTextLowerBound(0.5)
	menu.ResizeWindow(float32(width), float32(height))

	//2DO: sounds

	// label
	var label2 glmenu.Label
	menu.AddLabel(&label2, "Start")

	label2.Text.SetPosition(-(float32(menu.Width-margin) - label2.Text.Width), 0)
	label2.Text.SetColor(0.5, 0.5, 0.5, 1)
	label2.OnClick = func(label *glmenu.Label, xPos, yPos float64) (err error) {
		menu.Toggle()
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
	label2.AddShadow(3, 0, 0, 0)

	return
}
