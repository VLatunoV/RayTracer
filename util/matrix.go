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
	return (m.Cell[0][0]*(m.Cell[1][1]*m.Cell[2][2]-m.Cell[1][2]*m.Cell[2][1]) +
		m.Cell[0][1]*(m.Cell[1][2]*m.Cell[2][0]-m.Cell[1][0]*m.Cell[2][2]) +
		m.Cell[0][2]*(m.Cell[1][0]*m.Cell[2][1]-m.Cell[1][1]*m.Cell[2][0]))
}

// A^-1 = 1/det(A) * Adj(A)
// Adj(A) = (c_ij)^T
// c_ij = cofactor_ij(A)
func (m *Matrix3) Inverse() Matrix3 {
	d := m.Determinant()
	result := Matrix3{}
	if d.IsZero() {
		return result
	}
	var rows, cols [2]int
	var row, col int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			row = 0
			col = 0
			for k := 0; k < 3; k++ {
				if i != k {
					rows[row] = k
					row++
				}
				if j != k {
					cols[col] = k
					col++
				}
			}
			result.Cell[j][i] = (m.Cell[rows[0]][cols[0]]*m.Cell[rows[1]][cols[1]] - m.Cell[rows[1]][cols[0]]*m.Cell[rows[0]][cols[1]]) / d
			if (i+j)%2 == 1 {
				result.Cell[j][i] = -result.Cell[j][i]
			}
		}
	}
	return result
}
