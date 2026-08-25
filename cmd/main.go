package main

import (
	"fmt"
	linaltype "linalg/internal/linalType"
)

func main() {
	testData := [2]float64{1.2, 3.2}

	var testVector2D *linaltype.Vector2D = linaltype.NewVector2D(testData)

	var anotherVector *linaltype.Vector2D = linaltype.NewVector2D([2]float64{0, 0})

	fmt.Println(testVector2D.GetLength())

	var err = anotherVector.Normalize()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
