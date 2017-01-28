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
	r.Pos.X *= transform.Scale.X
	r.Pos.Y *= transform.Scale.Y
	r.Pos.Z *= transform.Scale.Z

	r.Pos.ApplyMatrix(transform.RotationMatrix)
	
	r.Pos.X += transform.Translate.X
	r.Pos.Y += transform.Translate.Y
	r.Pos.Z += transform.Translate.Z

	r.Dir.ApplyMatrix(transform.RotationMatrix)
}

func (r *Ray) Reverse(transform util.Transform) {
	r.Pos.X -= transform.Translate.X
	r.Pos.Y -= transform.Translate.Y
	r.Pos.Z -= transform.Translate.Z

	r.Pos.ApplyMatrix(transform.InverseRotation)

	r.Pos.X /= transform.Scale.X
	r.Pos.Y /= transform.Scale.Y
	r.Pos.Z /= transform.Scale.Z

	r.Dir.ApplyMatrix(transform.InverseRotation)
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
