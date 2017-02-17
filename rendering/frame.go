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

func (f *Frame) SetPixel(x, y int, r, g, b byte) {
	index := x + y*f.Width
	f.Data[index].R = r
	f.Data[index].G = g
	f.Data[index].B = b
}
