package linaltype

type Matrix struct {
	data [][]float64
}

func NewMatrix(rows int, cols int, data []float64) *Matrix {

	var matrix [][]float64 = make([][]float64, rows)

	for i := range matrix {
		matrix[i] = make([]float64, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = data[i*cols+j]
		}
	}

	return &Matrix{
		data: matrix,
	}
}
