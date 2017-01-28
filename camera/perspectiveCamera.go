package camera

import (
	"github.com/VLatunoV/RayTracer/util"
)

type PerspectiveCamera struct {
	Transform   util.Transform
	AspectRatio util.Float
	FOV         util.Float
}

func MakePerspectiveCamera(aspectRatio, fov float64) PerspectiveCamera {
	return PerspectiveCamera{
		Transform: util.GetIdentityTransform(),
		AspectRatio: util.Float(aspectRatio),
		FOV: util.Float(fov),
	}
}
