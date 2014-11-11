package main

import (
	"fmt"
	glmenu "github.com/4ydx/glmenu"
	gltext "github.com/4ydx/gltext"
	glfw "github.com/go-gl/glfw3"
	gl32 "github.com/go-gl/glow/gl-core/3.2/gl"
	"github.com/go-gl/glow/gl-core/3.3/gl"
	"runtime"
)

var useStrictCoreProfile = (runtime.GOOS == "darwin")

func findCenter(windowWidth int, windowHeight int, x1, x2 gltext.Point) (lowerLeft gltext.Point) {
	widthHalf := windowWidth / 2
	heightHalf := windowHeight / 2

	lineWidthHalf := (x2.X - x1.X) / 2
	lineHeightHalf := (x2.Y - x1.Y) / 2

	lowerLeft.X = float32(widthHalf) - lineWidthHalf
	lowerLeft.Y = float32(heightHalf) - lineHeightHalf
	return
}

func errorCallback(err glfw.ErrorCode, desc string) {
	fmt.Printf("%v: %v\n", err, desc)
}

func keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyM && action == glfw.Press {
		menu.Toggle()
	}
}

func mouseButtonCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if button == glfw.MouseButtonLeft && action == glfw.Press {
		xPos, yPos := window.GetCursorPosition()
		//fmt.Println("button", button, xPos, yPos)
		menu.ScreenClick(xPos, yPos)
	}
}

var menu glmenu.Menu
var window *glfw.Window

func main() {
	var err error

	runtime.LockOSThread()

	glfw.SetErrorCallback(errorCallback)
	if !glfw.Init() {
		panic("glfw error")
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	if useStrictCoreProfile {
		glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True)
		glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)
	}
	glfw.WindowHint(glfw.OpenglDebugContext, glfw.True)

	window, err = glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	window.SetKeyCallback(keyCallback)
	window.SetMouseButtonCallback(mouseButtonCallback)

	if err := gl.Init(); err != nil {
		panic(err)
	}
	if err := gl32.Init(); err != nil {
		fmt.Println("could not initialize GL 3.2")
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("Opengl version", version)

	width, height := window.GetSize()

	// menu settings
	menu.ResizeWindow(float32(width), float32(height))
	menu.SetDimension(400, 350)
	lowerLeft := menu.FindCenter()
	menu.Load(lowerLeft)

	// add text
	var text glmenu.Text
	scale := int32(32)
	err = text.Load(scale, 32, 127)
	if err != nil {
		panic(err)
	}
	x1, x2 := text.SetString("test me")
	center := findCenter(width, height, x1, x2)
	text.SetPosition(center.X, center.Y)
	text.SetTextLowerBound(0.5)
	text.SetColor(0, 0, 0, 1)
	text.ResizeWindow(float32(width), float32(height))
	text.OnClick = func(xPos, yPos float64) (err error) {
		fmt.Println("clicked at", xPos, yPos)
		return
	}
	menu.Text = append(menu.Text, text)

	gl.ClearColor(0, 0, 0, 0.0)
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		if menu.Draw() {
			// pause gameplay
		} else {
			// do stuff
		}
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
