package util

import (
	. "../texture"
	"math"
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
	uvs UV
}

func (n Float) isZero() bool {
	return math.Abs(n) < 1e-6
}

func (v *Vec3) Length() Float {
	return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
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

func generateRandomSample(low, high Float) Float {
	return 0
}

func generateRandomSampleInCircle(rad Float) (Float, Float) {
	return 0, 0
}

func dot(p, q Vec3) Float {
	p.Normalize()
	q.Normalize()
	return Float(p.X * q.X + p.Y * q.Y + p.Z * q.Z)
}
