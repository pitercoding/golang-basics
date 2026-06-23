package main

import "fmt"

func main() {
	numbers := []int{5, 2, 4, 6, 1, 3}

	fmt.Println("Original Slice:", numbers)

	result := insertionSort(numbers)

	fmt.Println("Ordered Slice:", result)
}

func insertionSort(numbers []int) []int {
	for i := 1; i < len(numbers); i++ {
		key := numbers[i]
		j := i - 1

		for j >= 0 && numbers[j] > key {
			numbers[j+1] = numbers[j]
			j--
		}

		numbers[j+1] = key
	}

	return numbers
}
