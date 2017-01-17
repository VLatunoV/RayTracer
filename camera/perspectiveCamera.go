package camera

import (
	"../util"
)

type PerspectiveCamera struct {
	Transform util.Transform
	AspectRatio util.Float
}
