package geometry

import (
	"github.com/VLatunoV/RayTracer/util"
)

// A canonical sphere is that which is centered on the origin and with radius one.
type Sphere struct {
}

func (s *Sphere) Intersect(r Ray) *IntersectInfo {
	result := IntersectInfo{}
	result.Ray = r

	dot := -util.Dot(r.Dir, r.Pos)
	critical := dot*dot - r.Pos.LengthSqr() + 1
	if critical < 0 {
		return nil
	}
	critical = util.Sqrt(critical)
	t1 := dot - critical
	t2 := dot + critical
	if t1 > 0 && !t1.IsZero() {
		result.IntersectPoint = util.Add(util.Mult(r.Dir, t1), r.Pos)
		result.Normal = result.IntersectPoint
		result.Distance = t1
		return &result

	} else if t2 > 0 && !t2.IsZero() {
		result.IntersectPoint = util.Add(util.Mult(r.Dir, t2), r.Pos)
		result.Normal = result.IntersectPoint.Neg()
		result.Distance = t2
		return &result
	}
	return nil
}
