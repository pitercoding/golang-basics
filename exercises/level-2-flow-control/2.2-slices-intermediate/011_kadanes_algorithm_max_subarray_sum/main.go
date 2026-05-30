package main

import "fmt"

func main() {
	numbers := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	currentSum := numbers[0]
	maxSum := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if currentSum+numbers[i] > numbers[i] {
			currentSum += numbers[i]
		} else {
			currentSum = numbers[i]
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	fmt.Printf("Maximum subarray sum: %d\n", maxSum)
}