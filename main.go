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
		Translate: util.Vec3{0, 10, -20},
		Rotate: util.Vec3{
			X: 20,
		},
	})
	result.Lights = []light.PointLight{
		{
			Color: texture.RGB{
				R: 255, G: 50, B: 50,
			},
			Intensity: 10000.0,
			Position: util.Vec3{
				X: -50, Y: 80, Z: 15,
			},
		},
		{
			Color: texture.RGB{
				R: 50, G: 50, B: 255,
			},
			Intensity: 10000.0,
			Position: util.Vec3{
				X: 50, Y: 75, Z: 0,
			},
		},
		{
			Color: texture.RGB{
				R: 255, G: 255, B: 255,
			},
			Intensity: 50.0,
			Position: util.Vec3{
				X: -2, Y: 15, Z: 0,
			},
		},
	}
	sphere := geometry.Sphere{}
	sphere2 := geometry.Sphere{}
	plane := geometry.Plane{}
	result.Nodes = []scene.Node{
		{
			Geometry:  &sphere,
			Transform: util.GetIdentityTransform(),
		},
		{
			Geometry:  &sphere2,
			Transform: util.GetIdentityTransform(),
		},
		{
			Geometry:  &plane,
			Transform: util.GetIdentityTransform(),
		},
	}
	result.Nodes[0].Transform.Translate = util.Vec3{-9, 7, 0}
	result.Nodes[0].Transform.Scale = util.Vec3{4, 4, 4}
	result.Nodes[1].Transform.Translate = util.Vec3{5, 4, 15}
	result.Nodes[1].Transform.Scale = util.Vec3{10, 10, 10}

	return &result
}
