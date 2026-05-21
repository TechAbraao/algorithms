package main

import (
	"fmt"
	"modules/algorithms"
)

func main() {
	var resultOne []int = algorithms.MoveZerosOne([]int{0, 1, 0, 3, 12})
	fmt.Println(resultOne)

	var resultTwo []int = algorithms.MoveZerosTwo([]int{0, 1, 0, 3, 12})
	fmt.Println(resultTwo)

}