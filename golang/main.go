package main

import (
	"fmt"
	"modules/algorithms"
)

func main() {
	listOne := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
	listTwo := []string{"cat", "dog", "bird", "fish"}
	listThree := []string{"go", "go", "go", "go", "python"}

	resultOne := algorithms.WordCounter(listOne)
	resultTwo := algorithms.WordCounter(listTwo)
	resultThree := algorithms.WordCounter(listThree)

	fmt.Println(
		"1: ", resultOne,
		"2: ", resultTwo,
		"3: ", resultThree,
	)

}