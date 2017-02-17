package util

import (
	"math"
)

type Float float64

const (
	Rad = math.Pi / 180.0
)

func (n Float) IsZero() bool {
	return math.Abs(float64(n)) < 1e-12
}

func (n Float) Equals(other Float) bool {
	return (n - other).IsZero()
}

func DegreeToRad(degree Float) Float {
	return degree * Rad
}

func GetSinCos(degree Float) (Float, Float) {
	rad := DegreeToRad(degree)
	return Float(Sin(rad)), Float(Cos(rad))
}

func Min(x, y Float) Float {
	if x > y {
		return y
	}
	return x
}

func Max(x, y Float) Float {
	if x < y {
		return y
	}
	return x
}

func Clamp(val, lower, upper Float) Float {
	return Max(Min(val, upper), lower)
}

func Sqrt(val Float) Float {
	return Float(math.Sqrt(float64(val)))
}

func Sin(val Float) Float {
	return Float(math.Sin(float64(val)))
}

func Cos(val Float) Float {
	return Float(math.Cos(float64(val)))
}

func Tan(val Float) Float {
	return Float(math.Tan(float64(val)))
}
