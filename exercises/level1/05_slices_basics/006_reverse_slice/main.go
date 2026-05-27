package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("Original slice:", numbers)

	left := 0
	right := len(numbers) - 1

	for left < right {
		numbers[left], numbers[right] = numbers[right], numbers[left]

		left++
		right--
	}

	fmt.Println("Reversed slice:", numbers)
}