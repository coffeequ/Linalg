package main

import (
	"fmt"
	linaltype "linalg/internal/linalType"
)

func main() {
	testData := [2]float64{11, 12}

	var testVector2D *linaltype.Vector2D = linaltype.NewVector2D(testData)

	var anotherVector *linaltype.Vector2D = linaltype.NewVector2D([2]float64{0, 0})

	fmt.Println(testVector2D.GetLength())

	anotherVector.SetX(3)

	anotherVector.SetY(2)

	testVector2D.MultpValue(2)

	testVector2D.Print()

	testVector2D.DivisionValue(2)

	testVector2D.Print()

	var err = anotherVector.Normalize()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
