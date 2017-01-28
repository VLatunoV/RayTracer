package rendering

import "github.com/VLatunoV/RayTracer/texture"

// Frame is a rectangle fragment of an image.
type Frame struct {
	X, Y          int
	Width, Height int
	Data          []texture.RGB
}

// MakeFrame creates an empty Frame.
func MakeFrame(xoff, yoff, width, height int) Frame {
	return Frame{
		X:      xoff,
		Y:      yoff,
		Width:  width,
		Height: height,
		Data:   make([]texture.RGB, width*height),
	}
}
