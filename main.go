package main

import (
	"runtime"

	"github.com/VLatunoV/RayTracer/camera"
	"github.com/VLatunoV/RayTracer/geometry"
	"github.com/VLatunoV/RayTracer/light"
	. "github.com/VLatunoV/RayTracer/rendering"
	"github.com/VLatunoV/RayTracer/scene"
	"github.com/VLatunoV/RayTracer/texture"
	"github.com/VLatunoV/RayTracer/util"
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

	renderer.Scene = makeMyScene(width, height)
	// renderer.ReadInput(file)
	renderer.Attach(visualizer)
	renderer.Render()
	visualizer.Run()
}

func makeMyScene(w, h int) *scene.Scene {
	result := scene.Scene{}
	result.Camera = camera.MakePerspectiveCamera(float64(w)/float64(h), 90)
	result.Camera.SetTransform(util.Transform{
		Translate: util.Vec3{0, 1, -20},
	})
	result.Lights = make([]light.Light, 1)
	result.Lights[0] = light.PointLight{
		Color: texture.RGB{
			R: 255, G: 255, B: 255,
		},
		Intensity: 1000.0,
		Position: util.Vec3{
			X: 0, Y: 20, Z: 0,
		},
	}
	result.Nodes = make([]scene.Node, 3)
	sphere := geometry.Sphere{}
	sphere2 := geometry.Sphere{}
	plane := geometry.Plane{}
	result.Nodes[0] = scene.Node{
		Geometry: &sphere,
		Transform: util.GetIdentityTransform(),
	}
	result.Nodes[0].Transform.Translate = util.Vec3{-4, 0, -7}
	result.Nodes[0].Transform.Scale = util.Vec3{4, 4, 4}
	result.Nodes[1] = scene.Node{
		Geometry: &sphere2,
		Transform: util.GetIdentityTransform(),
	}
	result.Nodes[1].Transform.Translate = util.Vec3{5, 4, 15}
	result.Nodes[1].Transform.Scale = util.Vec3{10, 10, 10}
	result.Nodes[2] = scene.Node{
		Geometry: &plane,
		Transform: util.GetIdentityTransform(),
	}

	return &result
}
