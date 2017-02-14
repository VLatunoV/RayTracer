package geometry

import (
	"testing"
	"github.com/VLatunoV/RayTracer/util"
)

func TestSphere_Intersect(t *testing.T) {
	r := Ray{
		Dir: util.Vec3{3, 1, 0},
	}
	s := Sphere{
		Transform: util.GetIdentityTransform(),
	}
	s.Transform.Translate = util.Vec3{7, 1, 0}
	s.Transform.Scale = util.Vec3{2, 2, 2}
	r.Dir.Normalize()
	ii, ok := s.Intersect(r)
	if !ok {
		t.Error("Expected ray to intersect the sphere, but it didn't.")
	}
	expectedIntersection := util.Vec3{X:5.130306154330093, Y:1.7101020514433642, Z:0}
	if !util.Sub(expectedIntersection, ii.IntersectPoint).LengthSqr().IsZero() {
		t.Errorf("Expected intersection point %+v, but was %+v", expectedIntersection, ii.IntersectPoint)
	}
	expectedNormal := util.Vec3{X:-0.9348469228349536, Y:0.3550510257216821, Z:0}
	if !util.Sub(expectedNormal, ii.Normal).LengthSqr().IsZero() {
		t.Errorf("Expected normal %+v, but was %+v", expectedNormal, ii.Normal)
	}
}
