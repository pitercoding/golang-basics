package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	k := 2

	result := rotateArray(numbers, k)

	fmt.Println("Original:", numbers)
	fmt.Println("Rotated:", result)
}

func rotateArray(numbers []int, k int) []int {
	n := len(numbers)

	if n == 0 {
		return numbers
	}

	k = k % n

	return append(
		numbers[n-k:],
		numbers[:n-k]...,
	)
}
