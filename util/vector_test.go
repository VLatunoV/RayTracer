package util

import (
	"fmt"
	"testing"
)

func TestDotProduct(t *testing.T) {
	pairs := [][2]Vec3{
		{Vec3{1, 0, 0}, Vec3{0, 1, 0}},
		{Vec3{1, 0, 0}, Vec3{0, 0, 1}},
		{Vec3{0, 1, 0}, Vec3{0, 0, 1}},

		{Vec3{1, 1, 0}, Vec3{0, 0, 1}},
		{Vec3{1, 1, 0}, Vec3{-1, 1, 0}},
		{Vec3{1, 1, 1}, Vec3{-1, -1, -1}},
	}
	expected := []Float{
		0,
		0,
		0,
		0,
		0,
		-1,
	}
	for i, p := range pairs {
		t.Run(fmt.Sprintf("Running dot product tests %d", i), func(tt *testing.T) {
			d := Dot(p[0].Normalized(), p[1].Normalized())
			if !d.Equals(expected[i]) {
				tt.Errorf("Expected %f, but was %f", expected[i], d)
			}
		})
	}
}
