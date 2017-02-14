package rendering

import "github.com/VLatunoV/RayTracer/scene"

type Renderer struct {
	ResolutionX int
	ResolutionY int

	requests  <-chan Frame
	completed chan Frame

	Scene *scene.Scene
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
}

func MakeRenderer(width, height int) *Renderer {
	return &Renderer{
		ResolutionX: width,
		ResolutionY: height,
		completed:   make(chan Frame),
	}
}
