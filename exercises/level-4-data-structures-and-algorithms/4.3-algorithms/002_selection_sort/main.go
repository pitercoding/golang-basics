package main

import "fmt"

func main() {
	numbers := []int{64, 25, 12, 22, 11}

	fmt.Println("Original Slice:", numbers)

	result := selectionSort(numbers)

	fmt.Println("Ordered Slice:", result)
}

func selectionSort(numbers []int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] < numbers[minIndex] {
				minIndex = j
			}
		}
		numbers[i], numbers[minIndex] = numbers[minIndex], numbers[i]
	}
	return numbers
}
