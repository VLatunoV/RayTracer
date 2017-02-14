package geometry

import "github.com/VLatunoV/RayTracer/util"

type Ray struct {
	Pos util.Vec3
	Dir util.Vec3
}

func Copy(ray *Ray) Ray {
	return Ray{
		Pos: ray.Pos,
		Dir: ray.Dir,
	}
}

func (r *Ray) Apply(transform util.Transform) {
	r.Pos.ApplyTransform(transform)

	r.Dir.X *= transform.Scale.X
	r.Dir.Y *= transform.Scale.Y
	r.Dir.Z *= transform.Scale.Z

	r.Dir.ApplyMatrix(transform.RotationMatrix)
}

func (r *Ray) Reverse(transform util.Transform) {
	r.Pos.ReverseTransform(transform)

	r.Dir.ApplyMatrix(transform.InverseRotation)

	r.Dir.X /= transform.Scale.X
	r.Dir.Y /= transform.Scale.Y
	r.Dir.Z /= transform.Scale.Z
}

func (r *Ray) Transformed(transform util.Transform) Ray {
	result := Copy(r)
	result.Apply(transform)
	return result
}

func (r *Ray) Reversed(transform util.Transform) Ray {
	result := Copy(r)
	result.Reverse(transform)
	return result
}
