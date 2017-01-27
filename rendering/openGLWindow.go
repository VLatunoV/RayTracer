package rendering

import (
	"github.com/go-gl/gl/v4.5-compatibility/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"unsafe"
)

type OpenGLWindow struct {
	width  int
	height int
	Name   string
	window *glfw.Window
}

func (w *OpenGLWindow) GetInputStream() <-chan Frame {
	return nil
}

func (w *OpenGLWindow) SetOutputStream(<-chan Frame) {

}

// Run should only be called from the main thread.
func (w *OpenGLWindow) Run() {
	// Must be called before any GLFW functions.
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	w.window, err = glfw.CreateWindow(w.width, w.height, w.Name, nil, nil)
	w.window.SetSizeLimits(w.width, w.height, w.width, w.height)
	if err != nil {
		panic(err)
	}

	w.window.MakeContextCurrent()

	// Must be called after MakeContextCurrent.
	err = gl.Init()
	if err != nil {
		panic(err)
	}
	defer gl.Finish()

	var image uint32
	var imageData unsafe.Pointer
	pixelData := make([]byte, w.width*w.height*4)
	for j := 0; j < w.height; j++ {
		for i := 0; i < w.width; i++ {
			pixelData[(i+j*w.width)*4] = byte(i * 255 / w.width)
			pixelData[(i+j*w.width)*4+1] = 0
			pixelData[(i+j*w.width)*4+2] = byte(j * 255 / w.height)
			pixelData[(i+j*w.width)*4+3] = 255
		}
	}
	imageData = unsafe.Pointer(&pixelData[0])

	// Generate a texture to draw on the window.
	gl.GenBuffers(1, &image)
	gl.BindTexture(gl.TEXTURE_2D, image)
	gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(w.width), int32(w.height), 0, gl.RGBA, gl.UNSIGNED_BYTE, imageData)

	// Setup draw params.
	gl.Ortho(0, float64(w.width), float64(w.height), 0, 0, 1)
	gl.Enable(gl.TEXTURE_2D)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	for !w.window.ShouldClose() {
		// Do OpenGL stuff.
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		w.drawTexture(image)
		w.handleEvents()
	}
}

func (w *OpenGLWindow) drawTexture(texture uint32) {
	gl.Color3f(1, 1, 1)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.Begin(gl.QUADS)
	gl.TexCoord2f(0, 0)
	gl.Vertex2i(0, 0)
	gl.TexCoord2f(0, 1)
	gl.Vertex2i(0, int32(w.height))
	gl.TexCoord2f(1, 1)
	gl.Vertex2i(int32(w.width), int32(w.height))
	gl.TexCoord2f(1, 0)
	gl.Vertex2i(int32(w.width), 0)
	gl.End()
}

func (w *OpenGLWindow) updateTexture(texture uint32, frame Frame) {
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexSubImage2D(gl.TEXTURE_2D, 0, int32(frame.X), int32(frame.Y), int32(frame.Width), int32(frame.Height), gl.RGB, gl.UNSIGNED_BYTE, unsafe.Pointer(&frame.Data[0]))
}

func (w *OpenGLWindow) handleEvents() {
	w.window.SwapBuffers()
	glfw.PollEvents()
	if w.window.GetKey(glfw.KeyEscape) == 1 {
		w.window.SetShouldClose(true)
	}
}

func MakeOpenGLWindow(width, height int) *OpenGLWindow {
	return &OpenGLWindow{
		width:  width,
		height: height,
		Name:   "RayTracer",
	}
}
