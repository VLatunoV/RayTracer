package geometry

import (
	"github.com/VLatunoV/RayTracer/texture"
	"github.com/VLatunoV/RayTracer/util"
)

type IntersectInfo struct {
	IntersectPoint util.Vec3
	Normal         util.Vec3
	UVS            texture.UV
	Ray            Ray
}

func (i *IntersectInfo) Apply(t util.Transform) {
	i.IntersectPoint.ApplyTransform(t)
	i.Normal.ApplyMatrix(t.TransposeInverse)
	//i.Normal.ApplyTransform(t)
	i.Normal.Normalize()
}

type Intersectable interface {
	// Intersect returns the geometry intersection info with a given normalized Ray.
	Intersect(Ray) (IntersectInfo, bool)
}
//X:0.9486832980505138 Y:0.31622776601683794 Z:0
//X:-0.9348469228349536 Y:0.3550510257216821 Z:0