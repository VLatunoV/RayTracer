package geometry

import (
	"github.com/VLatunoV/RayTracer/util"
	"github.com/VLatunoV/RayTracer/texture"
)

type IntersectInfo struct {
	InterPoint util.Vec3
	Normal     util.Vec3
	UVS        texture.UV
}

type Intersectable interface {
	Intersect(*IntersectInfo)
}
