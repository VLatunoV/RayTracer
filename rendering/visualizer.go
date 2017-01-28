package rendering

type Visualizer interface {
	// SetOutputStream takes a read only Frame channel and visualizes the Frames from it.
	SetOutputStream(<-chan Frame)

	// GetInputStream returns a read only Frame channel that is used to pass Frame requests to the renderer.
	GetInputStream() <-chan Frame

	// Run starts the visualizer. The Visualizer must cleanup before Run exists.
	Run()
}
