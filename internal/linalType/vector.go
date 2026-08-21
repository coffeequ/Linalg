package linaltype

import (
	"fmt"
)

type Vector struct {
	Data [2]float64
}

// func (имя_параметра тип_получателя) име_метода(параметры) (тип возвращаемого значения)
func New(data [2]float64) *Vector {
	return &Vector{Data: data}
}

func (curr *Vector) Add(data [2]float64) {
	curr.Data[0] += data[0]
	curr.Data[1] += data[1]
}

func (curr Vector) Getter() [2]float64 {
	return curr.Data
}

func (curr Vector) Print() {
	fmt.Println(curr.Getter())
}
