package linaltype

import (
	"fmt"
	"math"
)

type Vector2D struct {
	data [2]float64
}

// func (имя_параметра тип_получателя) име_метода(параметры) (тип возвращаемого значения)
func NewVector2D(data [2]float64) *Vector2D {
	return &Vector2D{data: data}
}

func (curr *Vector2D) X() float64 {
	return curr.data[0]
}

func (curr *Vector2D) Y() float64 {
	return curr.data[1]
}

func (curr *Vector2D) Add(antoher *Vector2D) (ptrRes *Vector2D) {
	var temp [2]float64

	temp[0] = curr.X() + antoher.X()

	temp[1] = curr.Y() + antoher.Y()

	ptrRes = &Vector2D{
		data: temp,
	}

	return
}

func (curr *Vector2D) Getter() [2]float64 {
	return curr.data
}

func (curr *Vector2D) Print() {
	fmt.Println(curr.Getter())
}

func (curr Vector2D) GetLength() float64 {
	return math.Sqrt(curr.data[0]*curr.data[0] + curr.data[1]*curr.data[1])
}

func (curr *Vector2D) Dot(antoher []float64) float64 {
	return curr.data[0]*antoher[0] + curr.data[1]*antoher[1]
}
