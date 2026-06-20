package main

import "fmt"

func main() {
	numbers := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	result := flatten(numbers)

	fmt.Println("Nested:", numbers)
	fmt.Println("Flatten:", result)
}

func flatten(numbers [][]int) []int {
	result := []int{}

	for _, nestedSlice := range numbers {
		result = append(result, nestedSlice...)
	}

	return result
}