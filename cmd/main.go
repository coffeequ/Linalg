package main

import (
	"fmt"
	linaltype "linalg/internal/linalType"
)

func main() {
	//testData := []float32{0, 0, 1, 2, 3}

	var tryUse, _ = linaltype.NewMatrix(2, 2, []float32{1, 2})

	var anotherUse, _ = linaltype.NewMatrix(2, 2, []float32{3, 4})

	var resSumm, _ = linaltype.NewMatrix(1, 1, []float32{1, 1, 2})

	resSumm.Add(tryUse, anotherUse)

	fmt.Println(resSumm.GetDataAll())

}
