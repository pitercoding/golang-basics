package main

import "fmt"

func main() {
	numbers := []int{-5, 10, -3, 8, 0, 15, -1}

	onlyPositives := []int{}

	for _, n := range numbers {
		if n >= 0 {
			onlyPositives = append(onlyPositives, n)
		}
	}

	fmt.Println("Original slice: ", numbers)
	fmt.Println("Only positives from original slice: ", onlyPositives)
}