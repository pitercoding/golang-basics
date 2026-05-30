package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	if len(numbers) == 0 {
		fmt.Println("Empty slice")
		return
	}

	first := numbers[0]
	rest := numbers[1:]

	rotated := append(rest, first)

	fmt.Println("Rotated slice:", rotated)
}

