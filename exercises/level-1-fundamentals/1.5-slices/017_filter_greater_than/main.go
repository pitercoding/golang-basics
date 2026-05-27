package main

import "fmt"

func main() {
	numbers := []int{10, 25, 3, 40, 18, 7, 60}

	var threshold int

	fmt.Print("Enter threshold value: ")
	fmt.Scanln(&threshold)

	filtered := filterGreaterThan(numbers, threshold)

	fmt.Println("Filtered slice:", filtered)
}

func filterGreaterThan(slice []int, threshold int) []int {
	var result []int

	for _, n := range slice {
		if n > threshold {
			result = append(result, n)
		}
	}

	return result
}
