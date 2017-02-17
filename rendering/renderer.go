package rendering

import (
	"github.com/VLatunoV/RayTracer/scene"
	"github.com/VLatunoV/RayTracer/geometry"
	"github.com/VLatunoV/RayTracer/util"
	"github.com/VLatunoV/RayTracer/texture"
	"runtime"
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

func (r *Renderer) Attach(v Visualizer) {
	v.SetInputStream(r.completed)
	r.requests = v.GetRequestStream()
}

// Render starts the renderer. It will not block.
// Panics if Attach was not called or no Scene was given.
func (r *Renderer) Render() {
	if r.requests == nil {
		panic("No visualizer attached.")
	}
	if r.Scene == nil {
		panic("No scene to render.")
	}
	numThreads := runtime.NumCPU() - 1
	for i := 0; i < numThreads; i++ {
		go r.handleRequests()
	}
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
				request.SetPixel(x, y, result.R, result.G, result.B)
			}
		}
		r.completed <- request
	}
}

func (r *Renderer) traceRay(ray geometry.Ray) texture.RGB {
	var closestIntersection *geometry.IntersectInfo
	hasIntersection := false
	for _, node := range r.Scene.Nodes {
		if ii := node.Intersect(ray); ii != nil {
			if !hasIntersection {
				hasIntersection = true
				closestIntersection = ii
			}
			if ii.Distance < closestIntersection.Distance {
				closestIntersection = ii
			}
		}
	}
	if closestIntersection == nil {
		return texture.RGB{0, 0, 0}
	}
	val := byte(255.0 / closestIntersection.Distance)
	return texture.RGB{val, val, val}
}

func (r *Renderer) shadowRay(start, end util.Vec3) {

}
