package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 6, 8, 10}
	target := 10

	left, right, found := twoSum(numbers, target)

	fmt.Println("\n=== Two Sum ===")
	fmt.Println("Slice:", numbers)
	fmt.Println("Target:", target)

	if found {
		fmt.Printf(
			"Found at index %d (%d) and %d (%d)\n",
			left,
			numbers[left],
			right,
			numbers[right],
		)
	} else {
		fmt.Println("Not found.")
	}
}

func twoSum(numbers []int, target int) (int, int, bool) {
	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == target {
			return left, right, true
		}

		if sum < target {
			left++
		}

		if sum > target {
			right--
		}

	}

	return -1, -1, false
}
