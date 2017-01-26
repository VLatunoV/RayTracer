package rendering

import "github.com/VLatunoV/RayTracer/texture"

// Frame is a rectangle fragment of an image
type Frame struct {
	X, Y int
	Width, Height int
	Data []texture.RGB
}
