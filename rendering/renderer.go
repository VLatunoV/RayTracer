package rendering

type Renderer struct {
	ResolutionX int
	ResolutionY int
}

func (r *Renderer) Attach(v Visualizer) {

}

// Render starts the renderer. It will not block.
func (r *Renderer) Render() {

}

func MakeRenderer(width, height int) *Renderer {
	return &Renderer{
		ResolutionX: width,
		ResolutionY: height,
	}
}
