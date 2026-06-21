package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	mapped := mapSquare(numbers)
	total := reduceSum(mapped)

	fmt.Println("Original:", numbers)
	fmt.Println("Mapped:", mapped)
	fmt.Println("Reduced:", total)
}

func mapSquare(numbers []int) []int {
	result := make([]int, 0, len(numbers))

	for _, number := range numbers {
		result = append(result, number*number)
	}

	return result
}

func reduceSum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}
