package rendering

type Visualizer interface {
	// SetInputStream takes a read only Frame channel and visualizes the Frames from it.
	SetInputStream(<-chan Frame)

	// GetRequestStream returns a read only Frame channel that is used to pass Frame requests to the renderer.
	GetRequestStream() <-chan Frame

	// Run starts the visualizer. It must cleanup before Run exists.
	Run()
}
