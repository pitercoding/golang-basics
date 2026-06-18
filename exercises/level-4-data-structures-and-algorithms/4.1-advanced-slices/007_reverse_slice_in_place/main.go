package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Original Slice:", numbers)
	reverseSlice(numbers)
	fmt.Println("Reversed Slice:", numbers)
}

func reverseSlice(numbers []int) {
	left := 0
	right := len(numbers) - 1

	for left < right {
		numbers[left], numbers[right] = numbers[right], numbers[left]

		left++
		right--
	}
}