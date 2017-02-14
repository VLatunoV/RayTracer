package scene

import (
	"github.com/VLatunoV/RayTracer/geometry"
	"github.com/VLatunoV/RayTracer/texture"
)

type Node struct {
	Geometry geometry.Intersectable
	Texture  texture.Shadeable
}
