package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	result := isSorted(numbers)

	fmt.Println("Slice:", numbers)
	fmt.Println("Sorted:", result)
}

func isSorted(numbers []int) bool {

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			return false
		}
	}

	return true
}
