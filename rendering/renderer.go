package rendering

type Renderer struct {

}

func (r *Renderer) Attach(v Visualizer) {

}

// Render starts the renderer. It will not block.
func (r *Renderer) Render() {

}

func MakeRenderer(width, height int) *Renderer {
	return &Renderer{}
}
