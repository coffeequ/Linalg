package main

import (
	"fmt"
	"linalg/internal/linalType"
)

func main() {
	testData := [2]float64{1.2, 3.2}

	temp := linaltype.New(testData)

	fmt.Println(*temp)
}

