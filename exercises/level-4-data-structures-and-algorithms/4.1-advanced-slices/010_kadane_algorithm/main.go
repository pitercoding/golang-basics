package main

import "fmt"

func main() {
	numbers := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	sum, subarray := kadaneWithSubarray(numbers)

	fmt.Println("Array:", numbers)
	fmt.Println("Max sum:", sum)
	fmt.Println("Best subarray:", subarray)
}

func kadaneWithSubarray(numbers []int) (int, []int) {
	if len(numbers) == 0 {
		return 0, nil
	}

	currentSum := numbers[0]
	maxSum := numbers[0]

	start := 0
	bestStart := 0
	bestEnd := 0

	for i := 1; i < len(numbers); i++ {
		num := numbers[i]

		// decide restart or continue
		if currentSum+num < num {
			currentSum = num
			start = i
		} else {
			currentSum += num
		}

		// update best
		if currentSum > maxSum {
			maxSum = currentSum
			bestStart = start
			bestEnd = i
		}
	}

	return maxSum, numbers[bestStart : bestEnd+1]
}
