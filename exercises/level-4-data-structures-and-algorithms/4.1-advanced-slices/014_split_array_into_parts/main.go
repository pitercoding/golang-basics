package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	parts := 3

	result := splitArray(numbers, parts)

	fmt.Println("Original:", numbers)
	fmt.Println("Parts:", result)
}

func splitArray(numbers []int, parts int) [][]int {
	if parts <= 0 {
		return nil
	}

	n := len(numbers)

	size := n / parts
	remainder := n % parts

	result := make([][]int, 0, parts)

	start := 0

	for i := 0; i < parts; i++ {
		end := start + size

		if remainder > 0 {
			end++
			remainder--
		}

		if start > n {
			result = append(result, []int{})
			continue
		}

		if end > n {
			end = n
		}

		result = append(result, numbers[start:end])
		start = end
	}

	return result
}