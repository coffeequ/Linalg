package linaltype

type Matrix struct {
	data [][]float64
}

func NewMatrix(rows int, cols int, dRows []float64, dCols []float64) *Matrix {
	var matrix [][]float64 = make([][]float64, rows)

	for i := range matrix {
		matrix[i] = make([]float64, cols)
	}

	
}
