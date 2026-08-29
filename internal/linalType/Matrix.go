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
		return &Matrix{}, errors.New("Rows or cols < 0!")
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

func (curr *Matrix) GetDataAll() [][]float32 {
	return slices.Clone(curr.data)
}

func (curr *Matrix) GetCurrentItem(row int, col int) float32 {
	return curr.data[row][col]
}

func (curr *Matrix) Cols() int {
	return curr.col
}

func (curr *Matrix) Rows() int {
	return curr.row
}

func (curr *Matrix) Add(valA, valB *Matrix) error {

	if curr == nil {
		return errors.New("Matrix is nill")
	}

	if valA == nil || valB == nil {
		return errors.New("Matrix is nill")
	}

	if valA.row != valB.row || valA.col != valB.col {
		return errors.New("Error. Matrix is not equal")
	}

	var rows int = valA.row

	if rows == 0 {
		curr.data = nil
		return nil
	}

	var cols int = valA.col

	for i := 0; i < rows; i++ {

	}

	return nil
}
