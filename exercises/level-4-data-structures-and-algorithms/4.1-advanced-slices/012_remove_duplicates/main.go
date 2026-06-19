package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 7, 8, 8, 9, 10}

	result := removeDuplicates(numbers)

	fmt.Println()
	fmt.Println("Original Slice:", numbers)
	fmt.Println("Without Duplicates:", result)

}

func removeDuplicates(numbers []int) []int {
	seen := make(map[int]bool)

	result := make([]int, 0, len(numbers))

	for _, number := range numbers {
		if !seen[number] {
			result = append(result, number)
			seen[number] = true
		}
	}

	return result
}
