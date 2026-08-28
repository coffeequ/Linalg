package linaltype

import (
	"errors"
	"math"
	"strconv"
)

// Каждая позиция в массиве соответствует позиции координаты. 0 - X, 1 - Y, 2 - Z
type Vector3D struct {
	data [3]float64
}

func NewVector3D(value [3]float64) *Vector3D {
	return &Vector3D{
		data: value,
	}
}

func (curr *Vector3D) Getter() [3]float64 {
	return curr.data
}

func (curr *Vector3D) ToString() string {
	return strconv.FormatFloat(curr.data[0], 'f', 2, 64) + " " + strconv.FormatFloat(curr.data[1], 'f', 2, 64) + strconv.FormatFloat(curr.data[2], 'f', 2, 64)
}

func (curr *Vector3D) X() float64 {
	return curr.data[0]
}

func (curr *Vector3D) Y() float64 {
	return curr.data[1]
}

func (curr *Vector3D) Z() float64 {
	return curr.data[2]
}

func (curr *Vector3D) SetX(val float64) {
	curr.data[0] = val
}

func (curr *Vector3D) SetY(val float64) {
	curr.data[1] = val
}

func (curr *Vector3D) SetZ(val float64) {
	curr.data[2] = val
}

func (curr *Vector3D) Add(another *Vector3D) (ptrRes *Vector3D) {
	ptrRes = &Vector3D{
		data: [3]float64{
			curr.X() + another.X(),
			curr.Y() + another.Y(),
			curr.Z() + another.Z()},
	}
	return
}

func (curr *Vector3D) MultpValue(val float64) {
	curr.data[0] *= val

	curr.data[1] *= val

	curr.data[2] *= val
}

func (curr *Vector3D) Cross(another *Vector3D) (prtRes *Vector3D) {
	return &Vector3D{
		data: [3]float64{
			(curr.Y()*another.Z() - curr.Z()*another.Y()),
			-1 * (curr.X()*another.Z() - curr.Z()*another.X()),
			(curr.X()*another.Y() - curr.Y()*another.X()),
		},
	}
}

func (curr *Vector3D) GetLength() float64 {
	return math.Sqrt(curr.data[0]*curr.data[0] + curr.data[1]*curr.data[1] + curr.data[2]*curr.data[2])
}

func (curr *Vector3D) Normalize() error {
	if curr.GetLength() == 0 {
		return errors.New("Error. X Y Z equal 0")
	}
	var vectLength = curr.GetLength()

	curr.data[0] /= vectLength

	curr.data[1] /= vectLength

	curr.data[2] /= vectLength

	return nil
}
