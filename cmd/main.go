package main

import (
	linaltype "linalg/internal/linalType"
)

func main() {
	testData := []float64{11, 12}

	var res *linaltype.Matrix = linaltype.NewMatrix(2, 3, testData)

	println(res)

	// var testVector2D *linaltype.Vector2D = linaltype.NewVector2D(testData)

	// var anotherVector *linaltype.Vector2D = linaltype.NewVector2D([2]float64{0, 0})

	// var tempVect3D *linaltype.Vector3D = linaltype.NewVector3D([3]float64{0.2, 1.2, 333.2})

	// fmt.Println(testVector2D.GetLength())

	// var text string = "Hello world"

	// fmt.Println(text)

	// fmt.Println(tempVect3D.ToString())

	// anotherVector.SetX(3)

	// anotherVector.SetY(2)

	// testVector2D.MultpValue(2)

	// var errDivisin error = testVector2D.DivisionValue(2)

	// if errDivisin != nil {
	// 	fmt.Println(errDivisin.Error())
	// }

	// fmt.Println(testVector2D.ToString())

	// var err = anotherVector.Normalize()

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

}
