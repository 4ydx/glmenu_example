package main

import (
	"fmt"
	glmenu "github.com/4ydx/glmenu"
	glfw "github.com/go-gl/glfw3"
	gl32 "github.com/go-gl/glow/gl-core/3.2/gl"
	"github.com/go-gl/glow/gl-core/3.3/gl"
	"runtime"
)

var useStrictCoreProfile = (runtime.GOOS == "darwin")

func errorCallback(err glfw.ErrorCode, desc string) {
	fmt.Printf("%v: %v\n", err, desc)
}

func keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyM && action == glfw.Press {
		mainMenu.Toggle()
	}
	if key == glfw.KeyO && action == glfw.Press {
		optionMenu.Toggle()
	}
}

func mouseButtonCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if button == glfw.MouseButtonLeft && action == glfw.Press {
		xPos, yPos := window.GetCursorPosition()
		mainMenu.ScreenClick(xPos, yPos)
	}
}

var mainMenu glmenu.Menu
var optionMenu glmenu.Menu
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

	mainMenuInit(window)
	optionMenuInit(window)
	mainMenu.Toggle()

	gl.ClearColor(0, 0, 0, 0.0)
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		xPos, yPos := window.GetCursorPosition()
		mainMenu.ScreenHover(xPos, yPos)
		if mainMenu.Draw() || optionMenu.Draw() {
			// pause gameplay
		} else {
			// do stuff
		}
		window.SwapBuffers()
		glfw.PollEvents()
	}
	mainMenu.Release()
	optionMenu.Release()
}
