package main

import (
	"fmt"
	"slices"
)

func main() {
	var numbers []int
	var quantity int

	fmt.Println("\n=== Max Value Finder ===")

	fmt.Print("How many numbers do you want to add? ")
	fmt.Scanln(&quantity)

	for i := 0; i < quantity; i++ {
		var number int

		fmt.Printf("Enter number %d: ", i+1)
		fmt.Scanln(&number)
		numbers = append(numbers, number)
	}

	result := max(numbers)

	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Maximum value: %d\n", result)
}

func max(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxValue := slices.Max(nums)

	return maxValue
}
