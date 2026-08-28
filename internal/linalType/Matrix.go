package linaltype

import (
	"errors"
	"slices"
)

type Matrix struct {
	data [][]float32
	row  int
	col  int
}

func NewMatrix(rows int, cols int, data []float32) (*Matrix, error) {

	if rows < 0 || cols < 0 {
		return nil, errors.New("Rows or cols < 0!")
	}

	var matrix [][]float32 = make([][]float32, rows)

	for i := range matrix {
		matrix[i] = make([]float32, cols)
	}

	idx := 0

	for j := 0; j < cols; j++ {
		for i := 0; i < rows; i++ {
			if idx < len(data) {
				matrix[i][j] = data[idx]
				idx++
			}
		}
	}

	return &Matrix{
		data: matrix,
		row:  rows,
		col:  cols,
	}, nil
}

func (curr *Matrix) GetData() [][]float32 {
	return slices.Clone(curr.data)
}
