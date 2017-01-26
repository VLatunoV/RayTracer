package rendering

import (
	"github.com/go-gl/gl/v4.5-compatibility/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type OpenGLWindow struct {
	width int
	height int
	Name string
}

func (w *OpenGLWindow) GetInputStream() <-chan Frame {
	return nil
}

func (w *OpenGLWindow) SetOutputStream(<-chan Frame) {

}

// Run should only be called from the main thread.
func (w *OpenGLWindow) Run() {
	// Must be called before any GLFW functions
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(w.width, w.height, w.Name, nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	// Must be called after MakeContextCurrent
	err = gl.Init()
	if err != nil {
		panic(err)
	}
	defer gl.Finish()

	// gl.Viewport(0, 0, 200, 200)

	for !window.ShouldClose() {
		// Do OpenGL stuff.
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.Begin(gl.QUADS)
		gl.Color3f(0, 1, 1)
		gl.Vertex2i(0, 0)
		gl.Vertex2i(0, 100)
		gl.Vertex2i(100, 100)
		gl.Vertex2i(100, 0)
		gl.End()
		window.SwapBuffers()
		glfw.PollEvents()
		if window.GetKey(glfw.KeyEscape) == 1 {
			window.SetShouldClose(true)
		}
	}
}

func MakeOpenGLWindow(width, height int) Visualizer {
	return &OpenGLWindow{
		width: width,
		height: height,
		Name: "RayTracer",
	}
}
