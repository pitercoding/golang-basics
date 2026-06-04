package main

import "fmt"

func main() {
	numbers := []float64{10, 20, 30, 40, 50}

	result := average(numbers)

	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Average: %.2f\n", result)
}

func average(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}

	var sum float64

	for _, n := range nums {
		sum += n
	}

	return sum / float64(len(nums))
}