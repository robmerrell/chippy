package ui

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

// Display holds everything needed to manage windows and draw to the screen
type Display struct {
	window *glfw.Window
}

// NewDisplay creates a new window and initializes OpenGL
func NewDisplay(width, height, pixelScale int) (*Display, error) {
	err := glfw.Init()
	if err != nil {
		return nil, err
	}

	// set up the window
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	window, err := glfw.CreateWindow(width*pixelScale, height*pixelScale, "Chippy", nil, nil)
	if err != nil {
		return nil, err
	}

	if err := gl.Init(); err != nil {
		return nil, err
	}

	window.MakeContextCurrent()

	gl.Disable(gl.DEPTH_TEST)
	gl.ClearColor(0.0, 0.215, 0.361, 1.0)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	gl.Ortho(0, float64(width), float64(height), 0, 0, -1)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	return &Display{window: window}, nil
}

// Stop terminates drawing to the window
func (d *Display) Stop() {
	glfw.Terminate()
}

/*
func (d *Display) Draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.Color3f(1, 1, 1)
	x := 0.0
	y := 0.0
	size := 1.0
	gl.Rectd(x, y, x+size, y+size)

	d.window.SwapBuffers()
}

func (d *Display) Loop() {
	for !d.window.ShouldClose() {
		d.Draw()
		glfw.PollEvents()
	}
}
*/
