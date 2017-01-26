package camera

import (
	"github.com/VLatunoV/RayTracer/util"
)

type PerspectiveCamera struct {
	Transform util.Transform
	AspectRatio util.Float
}
