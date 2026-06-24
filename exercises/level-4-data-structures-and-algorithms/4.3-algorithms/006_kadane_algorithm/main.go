package main

import "fmt"

func main() {
	numbers := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	maxSum := kadane(numbers)

	fmt.Println("\n=== Kadane Algorithm ===")
	fmt.Println("Slice:", numbers)
	fmt.Println("Maximum Sum:", maxSum)
}

func kadane(numbers []int) int {
	currentSum := numbers[0]
	maxSum := numbers[0]

	for _, value := range numbers[1:] {
		if currentSum+value > value {
			currentSum += value
		} else {
			currentSum = value
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}
