package main

import "fmt"

func main() {
	numbers := []int{1, 4, 20, 3, 10, 5}
	target := 33

	left := 0
	currentSum := 0

	found := false

	for right := 0; right < len(numbers); right++ {
		currentSum += numbers[right]

		for currentSum > target {
			currentSum -= numbers[left]
			left++
		}

		if currentSum == target {
			fmt.Printf(
				"Subarray found from index %d to %d: %v\n",
				left,
				right,
				numbers[left:right+1],
			)

			found = true
			break
		}
	}

	if !found {
		fmt.Println("No subarray found")
	}
}
