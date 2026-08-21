package main

import (
	"linalg/internal/linalType"
)

func main() {
	testData := [2]float64{1.2, 3.2}

	temp := linaltype.New(testData)

	temp.Print()

	temp.Add(testData)

	temp.Print()
}

