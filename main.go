package main

import (
	"fmt"
	"github.com/4ydx/glmenu"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"runtime"
)

var useStrictCoreProfile = (runtime.GOOS == "darwin")

func keyCallback(
	w *glfw.Window,
	key glfw.Key,
	scancode int,
	action glfw.Action,
	mods glfw.ModifierKey,
) {
	if mainMenu.Visible && action == glfw.Release {
		if mods == glfw.ModShift {
			mainMenu.KeyRelease(key, true)
		} else {
			mainMenu.KeyRelease(key, false)
		}
	} else {
		if key == glfw.KeyM && action == glfw.Press {
			if optionMenu.Visible {
				optionMenu.Toggle()
			}
			mainMenu.Toggle()
		}
		if key == glfw.KeyO && action == glfw.Press {
			if !mainMenu.Visible {
				optionMenu.Toggle()
			}
		}
	}
}

func mouseButtonCallback(
	w *glfw.Window,
	button glfw.MouseButton,
	action glfw.Action,
	mods glfw.ModifierKey,
) {
	xPos, yPos := w.GetCursorPos()
	if button == glfw.MouseButtonLeft && action == glfw.Press {
		mainMenu.MouseClick(xPos, yPos, glmenu.MouseLeft)
		optionMenu.MouseClick(xPos, yPos, glmenu.MouseLeft)
	}
	if button == glfw.MouseButtonLeft && action == glfw.Release {
		mainMenu.MouseRelease(xPos, yPos, glmenu.MouseLeft)
		optionMenu.MouseRelease(xPos, yPos, glmenu.MouseLeft)
	}
}

var mainMenu glmenu.Menu
var optionMenu glmenu.Menu
var window *glfw.Window

func main() {
	var err error

	runtime.LockOSThread()

	err = glfw.Init()
	if err != nil {
		panic("glfw error")
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	if useStrictCoreProfile {
		glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
		glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	}
	glfw.WindowHint(glfw.OpenGLDebugContext, glfw.True)

	// fullscreen
	primary := glfw.GetPrimaryMonitor()
	vm := primary.GetVideoMode()
	w, h := vm.Width, vm.Height // you should probably pick one in another manner

	fmt.Println("w", w, "h", h)
	window, err = glfw.CreateWindow(w, h, "Testing", primary, nil)
	// fullscreen

	// windowed
	// window, err = glfw.CreateWindow(640, 480, "Testing", nil, nil)
	// windowed

	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	window.SetKeyCallback(keyCallback)
	window.SetMouseButtonCallback(mouseButtonCallback)

	if err := gl.Init(); err != nil {
		panic(err)
	}
	/*
		if err := gl32.Init(); err != nil {
			fmt.Println("could not initialize GL 3.2")
		}
	*/
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("Opengl version", version)

	mainMenuInit(window)
	optionMenuInit(window)
	mainMenu.Toggle()

	gl.ClearColor(0, 0, 0, 0.0)
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		xPos, yPos := window.GetCursorPos()
		mainMenu.MouseHover(xPos, yPos)
		optionMenu.MouseHover(xPos, yPos)
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
