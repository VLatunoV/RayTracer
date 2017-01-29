package util

type Transform struct {
	Translate Vec3
	Rotate    Vec3
	Scale     Vec3

	RotationMatrix  Matrix3
	InverseRotation Matrix3
}

func GetIdentityTransform() Transform {
	return Transform{
		Translate:       Vec3{0, 0, 0},
		Rotate:          Vec3{0, 0, 0},
		Scale:           Vec3{1, 1, 1},
		RotationMatrix:  GetIdentityMatrix(),
		InverseRotation: GetIdentityMatrix(),
	}
}

func (t *Transform) GetRotationMatrixX() Matrix3 {
	sin, cos := GetSinCos(t.Rotate.X)
	return Matrix3{
		Cell: [3][3]Float{
			{1, 0, 0},
			{0, cos, -sin},
			{0, sin, cos},
		},
	}
}

func (t *Transform) GetRotationMatrixY() Matrix3 {
	sin, cos := GetSinCos(t.Rotate.Y)
	return Matrix3{
		Cell: [3][3]Float{
			{cos, 0, sin},
			{0, 1, 0},
			{-sin, 0, cos},
		},
	}
}

func (t *Transform) GetRotationMatrixZ() Matrix3 {
	sin, cos := GetSinCos(t.Rotate.Z)
	return Matrix3{
		Cell: [3][3]Float{
			{cos, -sin, 0},
			{sin, cos, 0},
			{0, 0, 1},
		},
	}
}

// CalculateRotationMatrix creates the RotationMatrix and InverseRotation for the Transform. The order of rotation is Y -> X -> Z
func (t *Transform) CalculateRotationMatrix() {
	rotationX := t.GetRotationMatrixX()
	rotationY := t.GetRotationMatrixY()
	rotationZ := t.GetRotationMatrixZ()
	temp := MultMatrix(&rotationZ, &rotationX)
	t.RotationMatrix = MultMatrix(&temp, &rotationY)
	t.InverseRotation = t.RotationMatrix.Inverse()
}
