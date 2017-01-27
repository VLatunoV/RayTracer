package geometry

import "github.com/VLatunoV/RayTracer/util"

type Intersectable interface {
	Intersect(*util.IntersectInfo)
}
