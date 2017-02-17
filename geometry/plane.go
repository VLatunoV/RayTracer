package geometry

import "github.com/VLatunoV/RayTracer/util"

// A canonical plane is that which passes through the Ox and Oz axes
type Plane struct {
}

func (p *Plane) Intersect(ray Ray) *IntersectInfo {
	result := IntersectInfo{}
	if ray.Pos.Y * ray.Dir.Y > 0 {
		return nil
	}
	result.Distance = -ray.Pos.Y / ray.Dir.Y
	if result.Distance.IsZero() {
		return nil
	}
	result.IntersectPoint = util.Mult(ray.Dir, result.Distance)
	result.Normal.Y = 1.0
	if ray.Pos.Y < 0 {
		result.Normal.Y = -1.0
	}
	return &result
}
