package main

import (
	"fmt"
)

func menuInit() (err error) {
	// menu settings
	width, height := window.GetSize()

	// 2DO: set up like the gltext package so that this is internal...
	lowerLeft := menu.FindCenter()
	fontScale := int32(20)
	menu.Load(lowerLeft, 400, 350, fontScale)
	menu.ResizeWindow(float32(width), float32(height))

	menu.Font.SetTextLowerBound(0.5)

	label := menu.AddLabel("test me", 0, 0)

	label.OnClick = func(xPos, yPos float64) (err error) {
		fmt.Println("clicked at", xPos, yPos)
		return
	}

	label.OnHover = func(xPos, yPos float64) (err error) {
		label.Text.AddScale(menu.TextScaleRate)
		return
	}
	return
}
