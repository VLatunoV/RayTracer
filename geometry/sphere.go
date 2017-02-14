package geometry

import (
	"github.com/VLatunoV/RayTracer/util"
)

type Sphere struct {
	Transform util.Transform
}

func MakeSphere(pos util.Vec3, rad util.Float) Sphere {
	result := Sphere{
		Transform: util.GetIdentityTransform(),
	}
	result.Transform.Translate = pos
	result.Transform.Scale = util.Vec3{rad, rad, rad}
	return result
}

func (s *Sphere) Intersect(r Ray) (IntersectInfo, bool) {
	r.Reverse(s.Transform)
	//factor := r.Dir.Length()
	r.Dir.Normalize()
	result := IntersectInfo{}
	result.Ray = r

	dot := -util.Dot(r.Dir, r.Pos)
	critical := dot*dot - r.Pos.LengthSqr() + 1
	if critical < 0 {
		return IntersectInfo{}, false
	}
	critical = util.Sqrt(critical)
	t1 := dot - critical
	t2 := dot + critical
	if t1 > 0 && !t1.IsZero() {
		result.IntersectPoint = util.Add(util.Mult(r.Dir, t1), r.Pos)
		result.Normal = result.IntersectPoint
		result.Apply(s.Transform)
		// distance to intersection = t1 / factor
		return result, true

	} else if t2 > 0 && !t2.IsZero() {
		result.IntersectPoint = util.Add(util.Mult(r.Dir, t2), r.Pos)
		result.Normal = result.IntersectPoint
		result.Apply(s.Transform)
		// distance to intersection = t2 / factor
		return result, true
	}
	return IntersectInfo{}, false
}
