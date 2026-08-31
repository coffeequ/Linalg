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

	if valA == nil || valB == nil {
		return errors.New("Matrix one or two is nill")
	}

	var rowsA int = valA.Rows()

	var rowsB int = valB.Rows()

	var colsA int = valA.Cols()

	var colsB int = valB.Cols()

	if rowsA != rowsB || colsA != colsB {
		return errors.New("Error. Matrix is not equal")
	}

	if rowsA == 0 || rowsB == 0 {
		curr.data = nil
		return nil
	}

	if len(curr.data) != rowsA {
		curr.data = make([][]float32, rowsA)
	}

	for i := 0; i < rowsA; i++ {
		if len(curr.data[i]) != colsA {
			curr.data[i] = make([]float32, colsA)
		}
		for j := 0; j < colsA; j++ {
			curr.data[i][j] = valA.data[i][j] + valB.data[i][j]
		}
	}

	curr.row = rowsA

	curr.col = colsA

	return nil
}

func (curr *Matrix) Sub(valA, valB *Matrix) error {

	if valA == nil || valB == nil {
		return errors.New("Matrix one or two is nill")
	}

	var rowsA int = valA.Rows()

	var rowsB int = valB.Rows()

	var colsA int = valA.Cols()

	var colsB int = valB.Cols()

	if rowsA != rowsB || colsA != colsB {
		return errors.New("Error. Matrix is not equal")
	}

	if rowsA == 0 || rowsB == 0 {
		curr.data = nil
		return nil
	}

	if len(curr.data) != rowsA {
		curr.data = make([][]float32, rowsA)
	}

	for i := 0; i < rowsA; i++ {
		if len(curr.data[i]) != colsA {
			curr.data[i] = make([]float32, colsA)
		}
		for j := 0; j < colsA; j++ {
			curr.data[i][j] = valA.data[i][j] - valB.data[i][j]
		}
	}

	curr.row = rowsA

	curr.col = colsA

	return nil
}

func (curr *Matrix) Multiply(val float32) error {

	if curr.data == nil {
		curr.data = nil
		return nil
	}

	for i := 0; i < curr.Rows(); i++ {
		for j := 0; j < curr.Cols(); j++ {
			curr.data[i][j] *= val
		}
	}

	return nil
}

func (curr *Matrix) Divide(val float32) error {

	if val == 0 {
		return errors.New("Divide by zero is not work")
	}

	if curr.data == nil {
		curr.data = nil
		return nil
	}

	for i := 0; i < curr.Rows(); i++ {
		for j := 0; j < curr.Cols(); j++ {
			curr.data[i][j] /= val
		}
	}

	return nil
}

// Принадлежит к последующей оптимизации
func (curr *Matrix) MultiplyMatrix(valA, valB *Matrix) error {

	if valA == nil || valB == nil {
		return errors.New("Source matrix is nill")
	}

	if valA.Cols() != valB.Rows() {
		return errors.New("Matrix between isn`t soglasov")
	}

	var rows int = valA.Rows()

	var inner int = valA.Cols()

	var cols int = valB.Cols()

	if len(curr.data) != valA.Rows() {
		curr.data = make([][]float32, rows)
	}

	for i := 0; i < rows; i++ {
		if len(curr.data[i]) != cols {
			curr.data[i] = make([]float32, cols)
		} else {
			clear(curr.data[i])
		}
	}

	for i := 0; i < rows; i++ {
		for k := 0; k < inner; k++ {
			var a float32 = valA.data[i][k]
			var rowB []float32 = valB.data[k]
			var rowC []float32 = curr.data[i]
			for j := 0; j < cols; j++ {
				rowC[j] += a * rowB[j]
			}
		}
	}

	curr.row = rows

	curr.col = cols

	return nil
}
