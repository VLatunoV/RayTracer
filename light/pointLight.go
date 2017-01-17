package light

import (
	"../util"
	"../texture"
)

type PointLight struct {
	Position util.Vec3
	Intensity util.Float
	Color texture.RGB
}
