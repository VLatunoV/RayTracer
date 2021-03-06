package geometry

import (
	"github.com/VLatunoV/RayTracer/util"
	"testing"
)

func TestRayTransformAndReverse(t *testing.T) {
	transform := util.Transform{
		Scale:     util.Vec3{1, 1, 1},
		Translate: util.Vec3{10, 0, 0},
		Rotate:    util.Vec3{30, 60, 90},
	}
	transform.CalculateRotationMatrix()
	ray := Ray{
		Pos: util.Vec3{0, 0, 0},
		Dir: util.Vec3{1, 0, 0},
	}
	orig := Copy(&ray)

	ray.Apply(transform)
	ray.Reverse(transform)

	if util.VecNotEq(orig.Pos, ray.Pos) {
		t.Errorf("Expected original ray [%v] to equal TT^-1 ray [%v]", orig, ray)
	}
}

func TestRayRotateX(t *testing.T) {
	transform := util.GetIdentityTransform()
	transform.Rotate.X = 90
	transform.CalculateRotationMatrix()
	ray := Ray{
		Pos: util.Vec3{1, 0, 2},
		Dir: util.Vec3{0, 3, 0},
	}
	ray.Apply(transform)
	expected := util.Vec3{1, -2, 0}
	if util.VecNotEq(ray.Pos, expected) {
		t.Errorf("Expected ray position %v, but got %v", expected, ray.Pos)
	}
	expected = util.Vec3{0, 0, 3}
	if util.VecNotEq(ray.Dir, expected) {
		t.Errorf("Expected ray direction %v, but got %v", expected, ray.Dir)
	}
}

func TestRayRotateYZ(t *testing.T) {
	transform := util.GetIdentityTransform()
	transform.Rotate.Y = 90
	transform.Rotate.Z = 90
	transform.CalculateRotationMatrix()
	ray := Ray{
		Pos: util.Vec3{1, 0, 2},
		Dir: util.Vec3{0, 2, 0},
	}
	ray.Apply(transform)
	expected := util.Vec3{2, 1, 0}
	if util.VecNotEq(ray.Pos, expected) {
		t.Errorf("Expected ray position %v, but got %v", expected, ray.Pos)
	}
	expected = util.Vec3{0, 0, 2}
	if util.VecNotEq(ray.Dir, expected) {
		t.Errorf("Expected ray direction %v, but got %v", expected, ray.Dir)
	}
}
