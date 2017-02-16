package geometry

import (
	"github.com/VLatunoV/RayTracer/texture"
	"github.com/VLatunoV/RayTracer/util"
)

type IntersectInfo struct {
	IntersectPoint util.Vec3
	Normal         util.Vec3
	UVS            texture.UV
	Distance       util.Float
	Ray            Ray
}

func (i *IntersectInfo) Apply(t util.Transform) {
	i.IntersectPoint.ApplyTransform(t)
	i.Normal.ApplyMatrix(t.TransposeInverse)
	i.Normal.Normalize()
}

type Intersectable interface {
	// Intersect returns the geometry intersection info with a given normalized Ray.
	Intersect(Ray) (IntersectInfo, bool)
}
