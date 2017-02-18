package geometry

import (
	"github.com/VLatunoV/RayTracer/util"
	"testing"
)

func TestSphere_Intersect(t *testing.T) {
	r := Ray{
		Dir: util.Vec3{3, 1, 0},
	}
	transform := util.GetIdentityTransform()
	transform.Translate = util.Vec3{7, 1, 0}
	transform.Scale = util.Vec3{2, 2, 2}
	s := &Sphere{}
	r.Reverse(transform)
	r.Dir.Normalize()
	ii := s.Intersect(r)
	if ii == nil {
		t.Error("Expected ray to intersect the sphere, but it didn't.")
		return
	}
	ii.Apply(transform)
	expectedIntersection := util.Vec3{X: 5.130306154330093, Y: 1.7101020514433642, Z: 0}
	if !util.Sub(expectedIntersection, ii.IntersectPoint).LengthSqr().IsZero() {
		t.Errorf("Expected intersection point %+v, but was %+v", expectedIntersection, ii.IntersectPoint)
	}
	expectedNormal := util.Vec3{X: -0.9348469228349536, Y: 0.3550510257216821, Z: 0}
	if !util.Sub(expectedNormal, ii.Normal).LengthSqr().IsZero() {
		t.Errorf("Expected normal %+v, but was %+v", expectedNormal, ii.Normal)
	}
}
