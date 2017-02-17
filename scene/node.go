package scene

import (
	"github.com/VLatunoV/RayTracer/geometry"
	"github.com/VLatunoV/RayTracer/texture"
	"github.com/VLatunoV/RayTracer/util"
)

type Node struct {
	Geometry  geometry.Intersectable
	Transform util.Transform

	Texture texture.Shadeable
}

func (n *Node) Intersect(ray geometry.Ray) *geometry.IntersectInfo {
	ray.Reverse(n.Transform)
	factor := ray.Dir.Length()
	ray.Dir.Normalize()

	result := n.Geometry.Intersect(ray)
	if result == nil {
		return nil
	}
	result.Ray = ray
	result.Apply(n.Transform)
	result.Distance /= factor
	return result
}
