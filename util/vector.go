package util

type Vec3 struct {
	X, Y, Z Float
}

func VecEq(a, b Vec3) bool {
	sub := Sub(a, b)
	return sub.X.IsZero() && sub.Y.IsZero() && sub.Z.IsZero()
}

func VecNotEq(a, b Vec3) bool {
	return !VecEq(a, b)
}

/*
func (v *Vec3) RotateAroundX(degree Float) {
	var rad, sin, cos, y_old, z_old float64
	rad = float64(DegreeToRad(degree))
	sin = math.Sin(rad)
	cos = math.Cos(rad)
	y_old = float64(v.Y)
	z_old = float64(v.Z)
	v.Y = Float(y_old * cos - z_old * sin)
	v.Z = Float(y_old * sin + z_old * cos)
}

func (v *Vec3) RotateAroundY(degree Float) {
	var rad, sin, cos, x_old, z_old float64
	rad = float64(DegreeToRad(degree))
	sin = math.Sin(rad)
	cos = math.Cos(rad)
	x_old = float64(v.X)
	z_old = float64(v.Z)
	v.X = Float(x_old * cos + z_old * sin)
	v.Z = Float(z_old * cos - x_old * sin)
}

func (v *Vec3) RotateAroundZ(degree Float) {
	var rad, sin, cos, x_old, y_old float64
	rad = float64(DegreeToRad(degree))
	sin = math.Sin(rad)
	cos = math.Cos(rad)
	x_old = float64(v.X)
	y_old = float64(v.Y)
	v.X = Float(x_old * cos - y_old * sin)
	v.Y = Float(x_old * sin + y_old * cos)
}
*/
func (v *Vec3) ApplyMatrix(matrix Matrix3) {
	x := v.X
	y := v.Y
	z := v.Z
	v.X = matrix.Cell[0][0]*x + matrix.Cell[0][1]*y + matrix.Cell[0][2]*z
	v.Y = matrix.Cell[1][0]*x + matrix.Cell[1][1]*y + matrix.Cell[1][2]*z
	v.Z = matrix.Cell[2][0]*x + matrix.Cell[2][1]*y + matrix.Cell[2][2]*z
}

func (v *Vec3) ApplyTransform(transform Transform) {
	v.X *= transform.Scale.X
	v.Y *= transform.Scale.Y
	v.Z *= transform.Scale.Z

	v.ApplyMatrix(transform.RotationMatrix)

	v.X += transform.Translate.X
	v.Y += transform.Translate.Y
	v.Z += transform.Translate.Z
}

func (v *Vec3) ReverseTransform(transform Transform) {
	v.X -= transform.Translate.X
	v.Y -= transform.Translate.Y
	v.Z -= transform.Translate.Z

	v.ApplyMatrix(transform.InverseRotation)

	v.X /= transform.Scale.X
	v.Y /= transform.Scale.Y
	v.Z /= transform.Scale.Z
}

func (v Vec3) Length() Float {
	return Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vec3) LengthSqr() Float {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v *Vec3) Normalize() {
	l := v.Length()
	if l.IsZero() == false {
		v.X /= l
		v.Y /= l
		v.Z /= l
	}
}

func (v Vec3) Normalized() Vec3 {
	result := Vec3{
		v.X,
		v.Y,
		v.Z,
	}
	result.Normalize()
	return result
}

func (v Vec3) Neg() Vec3 {
	return Vec3{-v.X, -v.Y, -v.Z}
}

func Add(a, b Vec3) Vec3 {
	return Vec3{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
	}
}

func Sub(a, b Vec3) Vec3 {
	return Add(a, b.Neg())
}

func Mult(a Vec3, s Float) Vec3 {
	return Vec3{a.X * s, a.Y * s, a.Z * s}
}

func Dot(a, b Vec3) Float {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func Reflect(in, normal Vec3) Vec3 {
	d := Dot(in, normal)
	return Add(in, Mult(normal, -2*d))
}
