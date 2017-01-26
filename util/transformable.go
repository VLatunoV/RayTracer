package util

import "github.com/VLatunoV/RayTracer/util"

type Transformable interface {
	Transform(util.Transform)
	Reverse(util.Transform)
}