package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50, 60, 70}
	target := 70

	fmt.Println("\n=== Binary Search ===")
	fmt.Println("Slice:", numbers)

	index, found := binarySearch(numbers, target)

	if found {
		fmt.Printf("Value %d found at index %d.\n", target, index)
	} else {
		fmt.Printf("Value %d not found.\n", target)
	}
}

func binarySearch(numbers []int, target int) (int, bool) {
	left := 0
	right := len(numbers) - 1

	for left <= right {
		mid := (left + right) / 2

		if numbers[mid] == target {
			return mid, true
		}

		if target > numbers[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1, false
}
