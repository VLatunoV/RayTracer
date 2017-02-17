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
	// TODO: change the returned value to a SRGB
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
	var R, G, B util.Float
	for _, light := range r.Scene.Lights {
		difference := util.Sub(light.Position, closestIntersection.IntersectPoint)
		lambertFactor := util.Dot(closestIntersection.Normal, difference.Normalized())
		if lambertFactor < 0 {
			continue
		}
		if r.isVisible(closestIntersection.IntersectPoint, light.Position) {
			contribution := lambertFactor * light.Intensity / difference.LengthSqr()
			R += contribution * util.Float(light.Color.R)
			G += contribution * util.Float(light.Color.G)
			B += contribution * util.Float(light.Color.B)
		}
	}
	return texture.RGB{
		R: byte(util.Clamp(R, 0, 255)),
		G: byte(util.Clamp(G, 0, 255)),
		B: byte(util.Clamp(B, 0, 255)),
	}
}

// isVisible checks whether the end point is visible from the start point.
func (r *Renderer) isVisible(start, end util.Vec3) bool {
	ray := geometry.Ray{
		Pos: start,
		Dir: util.Sub(end, start),
	}
	distance := ray.Dir.Length()
	ray.Dir.Normalize()

	for _, node := range r.Scene.Nodes {
		if ii := node.Intersect(ray); ii != nil {
			if ii.Distance < distance {
				return false
			}
		}
	}
	return true
}
