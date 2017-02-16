package scene

import (
	"github.com/VLatunoV/RayTracer/geometry"
	"github.com/VLatunoV/RayTracer/texture"
	"github.com/VLatunoV/RayTracer/util"
)

type Node struct {
	Geometry  geometry.Intersectable
	Transform util.Transform

	Texture   texture.Shadeable
}

func (n *Node) Intersect(ray geometry.Ray) (geometry.IntersectInfo, bool) {
	ray.Reverse(n.Transform)
	factor := ray.Dir.Length()
	ray.Dir.Normalize()

	result, ok := n.Geometry.Intersect(ray)
	if !ok {
		return result, ok
	}
	result.Ray = ray
	result.Apply(n.Transform)
	result.Distance /= factor
	return result, ok
}
