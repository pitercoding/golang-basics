package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	if len(numbers) == 0 {
		fmt.Println("Empty slice")
		return
	}

	last := numbers[len(numbers)-1]
	rest := numbers[:len(numbers)-1]

	rotated := append([]int{last}, rest...)

	fmt.Println("Rotated slice:", rotated)
}
