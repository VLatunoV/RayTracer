package util

type Matrix3 struct {
	Cell [3][3]Float
}

func GetIdentityMatrix() Matrix3 {
	return Matrix3{
		Cell: [3][3]Float{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
	}
}

func MultMatrix(a, b *Matrix3) Matrix3 {
	result := Matrix3{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				result.Cell[i][j] += a.Cell[i][k] * b.Cell[k][j]
			}
		}
	}
	return result
}

func (m *Matrix3) Determinant() Float {
	return (m.Cell[0][0] * (m.Cell[1][1] * m.Cell[2][2] - m.Cell[1][2] * m.Cell[2][1]) +
	      	m.Cell[0][1] * (m.Cell[1][2] * m.Cell[2][0] - m.Cell[1][0] * m.Cell[2][2]) +
		m.Cell[0][2] * (m.Cell[1][0] * m.Cell[2][1] - m.Cell[1][1] * m.Cell[2][0]))
}

func (m *Matrix3) Inverse() Matrix3 {
	d := m.Determinant()
	result := Matrix3{}
	if d.IsZero() {
		return result
	}
	var rows, cols [2]int
	var crow, ccol int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			crow = 0
			ccol = 0
			for k := 0; k < 3; k++ {
				if i != k {
					rows[crow] = k
					crow++
				}
				if j != k {
					cols[ccol] = k
					ccol++
				}
			}
			result.Cell[i][j] = (m.Cell[rows[0]][cols[0]] * m.Cell[rows[1]][cols[1]] - m.Cell[rows[1]][cols[0]] * m.Cell[rows[0]][cols[1]]) / d
		}
	}
	return result
}
