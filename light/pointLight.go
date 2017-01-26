package light

import (
	"github.com/VLatunoV/RayTracer/util"
	"github.com/VLatunoV/RayTracer/texture"
)

type PointLight struct {
	Position util.Vec3
	Intensity util.Float
	Color texture.RGB
}
