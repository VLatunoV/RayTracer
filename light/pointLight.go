package light

import (
	"github.com/VLatunoV/RayTracer/texture"
	"github.com/VLatunoV/RayTracer/util"
)

type PointLight struct {
	Position  util.Vec3
	Intensity util.Float
	Color     texture.RGB
}
