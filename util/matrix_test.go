package util

import (
	"testing"
)

func eqMatrix(a, b *Matrix3) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !a.Cell[i][j].Equals(b.Cell[i][j]) {
				return false
			}
		}
	}
	return true
}

func notEqMatrix(a, b *Matrix3) bool {
	return !eqMatrix(a, b)
}

func TestMultMatrix(t *testing.T) {
	a := Matrix3{
		Cell: [3][3]Float{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
	}
	b := Matrix3{
		Cell: [3][3]Float{
			{9, 8, 7},
			{6, 5, 4},
			{3, 2, 1},
		},
	}
	c := MultMatrix(&a, &b)
	expected := Matrix3{
		Cell: [3][3]Float{
			{30, 24, 18},
			{84, 69, 54},
			{138, 114, 90},
		},
	}
	if notEqMatrix(&c, &expected) {
		t.Errorf("Expected matrix to equal %v, but was %v", expected, c)
	}
}

func TestInverseIdentity(t *testing.T) {
	I := GetIdentityMatrix()
	II := I.Inverse()
	d := I.Determinant()
	dd := II.Determinant()
	if !d.Equals(1) {
		t.Errorf("Expected determinant of identity to equal 1, but was %f", d)
	}
	if !dd.Equals(1) {
		t.Errorf("Expected determinant of inverse to equal 1, but was %f", d)
	}
	if notEqMatrix(&I, &II) {
		t.Errorf("Expected inverse of identity to equal the identity, but was %v", II)
	}
}

func TestInverseMatrix(t *testing.T) {
	a := Matrix3{
		Cell: [3][3]Float{
			{3, -5, 1},
			{2, 2, 8},
			{5, 1, 2},
		},
	}
	d := a.Determinant()
	if !d.Equals(-200) {
		t.Errorf("Expected determinant to equal -203, but was %f", d)
	}
	b := a.Inverse()
	expectedInverse := Matrix3{
		Cell: [3][3]Float{
			{0.020, -0.055, 0.210},
			{-0.180, -0.005, 0.110},
			{0.040, 0.140, -0.080},
		},
	}
	if notEqMatrix(&b, &expectedInverse) {
		t.Errorf("Expected inverse to equal %v, but was %v", expectedInverse, b)
	}
	ab := MultMatrix(&a, &b)
	ba := MultMatrix(&b, &a)
	expected := GetIdentityMatrix()
	if notEqMatrix(&expected, &ab) {
		t.Errorf("Expected AA^-1 to equal %v, but was %v", expected, ab)
	}
	if notEqMatrix(&expected, &ba) {
		t.Errorf("Expected A^-1A to equal %v, but was %v", expected, ba)
	}
}
