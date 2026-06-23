package main

import "fmt"

func main() {
	numbers := []int{5, 3, 8, 1, 2}

	fmt.Println("Original Slice:", numbers)

	result := bubbleSort(numbers)

	fmt.Println("Ordered Slice:", result)
}

func bubbleSort(numbers []int) []int {
	
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	return numbers
}