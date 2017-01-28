package util

import (
	"math"
)

type Float float64

func (n Float) IsZero() bool {
	return math.Abs(float64(n)) < 1e-12
}

func (n Float) Equals(other Float) bool {
	return (n - other).IsZero()
}

func GenerateRandomSample(low, high Float) Float {
	return 0
}

func GenerateRandomSampleInCircle(rad Float) (Float, Float) {
	return 0, 0
}

func DegreeToRad(degree Float) float64 {
	return float64(degree) * math.Pi / 180.0
}

func GetSinCos(degree Float) (Float, Float) {
	rad := DegreeToRad(degree)
	return Float(math.Sin(rad)), Float(math.Cos(rad))
}
