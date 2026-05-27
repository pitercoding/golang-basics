package main

import "fmt"

func main() {
	numbers := []int{5, 2, 9, 1, 7, 3, 8}

	min := numbers[0]
	max := numbers[0]

	for _, n := range numbers {
		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
