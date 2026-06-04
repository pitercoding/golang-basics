package main

import "fmt"

func main() {
	var numbers []int
	var quantity int

	fmt.Println("\n=== Min Value Finder ===")

	fmt.Print("How many numbers do you want to add? ")
	fmt.Scanln(&quantity)

	for i := 0; i < quantity; i++ {
		var number int

		fmt.Printf("Enter number %d: ", i+1)
		fmt.Scanln(&number)

		numbers = append(numbers, number)
	}

	result := min(numbers)

	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Min value: %d\n", result)
}

func min(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	minValue := nums[0]

	for _, n := range nums {
		if n < minValue {
			minValue = n
		}
	}

	return minValue
}