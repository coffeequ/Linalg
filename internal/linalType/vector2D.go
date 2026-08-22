package linaltype

import (
	"fmt"
	"math"
)

type Vector2D struct {
	Data [2]float64
}

// func (имя_параметра тип_получателя) име_метода(параметры) (тип возвращаемого значения)
func New(data [2]float64) *Vector2D {
	return &Vector2D{Data: data}
}

func (curr *Vector2D) Add(data [2]float64) {
	curr.Data[0] += data[0]
	curr.Data[1] += data[1]
}

func (curr Vector2D) Getter() [2]float64 {
	return curr.Data
}

func (curr Vector2D) Print() {
	fmt.Println(curr.Getter())
}

func (curr Vector2D) GetLength() float64 {
	return math.Sqrt(curr.Data[0]*curr.Data[0] + curr.Data[1]*curr.Data[1])
}

func (curr *Vector2D) Dot(antoher [2]float64) float64 {
	return curr.Data[0]*antoher[0] + curr.Data[1]*antoher[1]
}
