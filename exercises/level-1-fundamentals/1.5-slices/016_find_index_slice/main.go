package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}

	var value int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&value)

	// modern:
	// index := slices.Index(numbers, value)

	index := findIndex(numbers, value)

	if index != -1 {
		fmt.Printf("Value %d found at index %d\n", value, index)
	} else {
		fmt.Printf("Value %d not found\n", value)
	}
}

func findIndex(slice []int, target int) int {
	for i, n := range slice {
		if n == target {
			return i
		}
	}
	return -1
}
