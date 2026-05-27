package main

import "fmt"

func main() {
	original := []int{10, 20, 30, 40, 50}

	copied := make([]int, len(original))

	copy(copied, original)

	copied[0] = 999

	fmt.Println("Original:", original)
	fmt.Println("Copied:", copied)
}