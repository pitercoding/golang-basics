package main

import "fmt"

func main() {
	numbers := []int{2, 1, 5, 1, 3, 2}
	windowSize := 3

	currentSum := 0

	for i := 0; i < windowSize; i++ {
		currentSum += numbers[i]
	}

	maxSum := currentSum

	for i := windowSize; i < len(numbers); i++ {
		currentSum =
			currentSum -
				numbers[i-windowSize] +
				numbers[i]

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	fmt.Printf("Maximum window sum: %d\n", maxSum)
}