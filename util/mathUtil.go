package util

import (
	"math"
	"github.com/VLatunoV/RayTracer/texture"
)

type Float float32

type Vec3 struct {
	X, Y, Z Float
}

type Transform struct {
	Translate Vec3
	Rotate Vec3
	Scale Vec3
}

type IntersectInfo struct {
	InterPoint Vec3
	Normal Vec3
	uvs texture.UV
}

func (n Float) isZero() bool {
	return math.Abs(float64(n)) < 1e-6
}

func (n Float) equals(other Float) bool {
	return (n - other).isZero()
}

func (v *Vec3) Length() Float {
	return Float(math.Sqrt(float64(v.X * v.X + v.Y * v.Y + v.Z * v.Z)))
}

func (v *Vec3) Normalize() {
	l := v.Length()
	if l.isZero() == false {
		v.X /= l
		v.Y /= l
		v.Z /= l
	}
}

func (v *Vec3) Normalized() Vec3 {
	result := Vec3{
		v.X,
		v.Y,
		v.Z,
	}
	result.Normalize()
	return result
}

func (v *Vec3) Neg() Vec3 {
	return Vec3{-v.X, -v.Y, -v.Z}
}

func generateRandomSample(low, high Float) Float {
	return 0
}

func generateRandomSampleInCircle(rad Float) (Float, Float) {
	return 0, 0
}

func add(a, b Vec3) Vec3 {
	return Vec3{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
	}
}

func sub(a, b Vec3) Vec3 {
	return add(a, b.Neg())
}

func mult(a Vec3, s Float) Vec3 {
	return Vec3{a.X * s, a.Y * s, a.Z * s}
}

func dot(a, b Vec3) Float {
	a.Normalize()
	b.Normalize()
	return Float(a.X * b.X + a.Y * b.Y + a.Z * b.Z)
}

func reflect(in, normal Vec3) Vec3 {
	normal.Normalize()
	d := dot(in, normal)
	return add(in, mult(normal, -2 * d))
}
