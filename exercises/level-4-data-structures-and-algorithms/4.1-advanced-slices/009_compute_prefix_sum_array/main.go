package main

import "fmt"

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	result := prefixSum(numbers)

	fmt.Println()
	fmt.Println("Original:", numbers)
	fmt.Println("Result:", result)

}

func prefixSum(numbers []int) []int {
	result := make([]int, 0, len(numbers))
	sum := 0

	for _, number := range numbers {
		sum += number
		result = append(result, sum)
	}

	return result
}
