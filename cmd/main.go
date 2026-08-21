package main

import (
	"fmt"
)

func main() {
	var temp int = 4
	var p *int

	p = &temp

	fmt.Println(*p)

	var numbers [5]int = [5]int{1, 2, 3, 4, 5}

	var pNumber *[5]int = &numbers

	fmt.Println((*pNumber))
}

