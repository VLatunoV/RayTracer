package rendering

import (
	"github.com/VLatunoV/RayTracer/scene"
	"github.com/VLatunoV/RayTracer/geometry"
	"github.com/VLatunoV/RayTracer/util"
)

type Renderer struct {
	ResolutionX int
	ResolutionY int

	requests  <-chan Frame
	completed chan Frame

	Scene *scene.Scene
}

func MakeRenderer(width, height int) *Renderer {
	return &Renderer{
		ResolutionX: width,
		ResolutionY: height,
		completed:   make(chan Frame),
	}
}

// Render starts the renderer. It will not block.
// Panics if Attach was not called or no Scene was given.
func (r *Renderer) Attach(v Visualizer) {
	v.SetInputStream(r.completed)
	r.requests = v.GetRequestStream()
}

func (r *Renderer) Render() {
	if r.requests == nil {
		panic("No visualizer attached.")
	}
	if r.Scene == nil {
		panic("No scene to render.")
	}
	go r.handleRequests()
}

func (r *Renderer) handleRequests() {
	for {
		request, ok := <-r.requests
		if !ok {
			break
		}
		for y := 0; y < request.Height; y++ {
			for x := 0; x < request.Width; x++ {
				ray := r.Scene.Camera.GetRay(
					(util.Float(request.X + x) + 0.5) / util.Float(r.ResolutionX),
					(util.Float(request.Y + y) + 0.5) / util.Float(r.ResolutionY))
				result := r.traceRay(ray)
				if result {
					request.SetPixel(x, y, 255, 255, 255)
				} else {
					request.SetPixel(x, y, 0, 0, 0)
				}
			}
		}
		r.completed <- request
	}
}

func (r *Renderer) traceRay(ray geometry.Ray) bool {
	for _, node := range r.Scene.Nodes {
		if _, ok := node.Geometry.Intersect(ray); ok {
			return true
		}
	}
	return false
}
