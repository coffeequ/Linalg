package linaltype

import (
	"errors"
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

func (curr *Vector2D) SetX(value float64) {
	curr.data[0] = value
}

func (curr *Vector2D) SetY(value float64) {
	curr.data[1] = value
}

func (curr *Vector2D) Add(antoher *Vector2D) (ptrRes *Vector2D) {

	var temp Vector2D

	temp.data[0] = curr.X() + antoher.X()

	temp.data[1] = curr.Y() + antoher.Y()

	ptrRes = &temp

	return
}

func (curr *Vector2D) Subtraction(antoher *Vector2D) (ptrRes *Vector2D) {

	var temp Vector2D

	temp.data[0] = curr.X() + antoher.X()

	temp.data[1] = curr.Y() + antoher.Y()

	ptrRes = &temp

	return
}

func (curr *Vector2D) Getter() [2]float64 {
	return curr.data
}

func (curr *Vector2D) Print() {
	fmt.Println(curr.Getter())
}

func (curr *Vector2D) GetLength() float64 {
	return math.Sqrt(curr.data[0]*curr.data[0] + curr.data[1]*curr.data[1])
}

func (curr *Vector2D) Dot(antoher Vector2D) float64 {
	return curr.data[0]*antoher.data[0] + curr.data[1]*antoher.data[1]
}

func (curr *Vector2D) MultpValue(value float64) {
	curr.data[0] *= value
	curr.data[1] *= value
}

func (curr *Vector2D) DivisionValue(value float64) error {
	if value == 0 {
		return errors.New("Error. Devision by zero is dosent work")
	}

	curr.data[0] /= value

	curr.data[1] /= value

	return nil
}

func (curr *Vector2D) Normalize() error {
	if curr.GetLength() == 0 {
		return errors.New("Error. X and Y equal 0")
	}
	var vectLength = curr.GetLength()

	curr.data[0] /= vectLength

	curr.data[1] /= vectLength
	return nil
}
