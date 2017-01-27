package main

import (
	"runtime"

	. "github.com/VLatunoV/RayTracer/rendering"
)

func init() {
	// This is needed to arrange that main() runs on main thread
	// only when using GLFW.
	runtime.LockOSThread()
}

func main() {
	width := 640
	height := 480
	visualizer := MakeOpenGLWindow(width, height)
	renderer := MakeRenderer(width, height)

	// renderer.ReadInput(file)
	renderer.Attach(visualizer)
	renderer.Render()
	visualizer.Run()
}
