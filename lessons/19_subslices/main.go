package main

import "fmt"

func main() {
	original := []int{1, 2, 3, 4, 5}
	subSlice := original[1:4] // Elementos do índice 1 a 3
	fmt.Println("Original:", original)
	fmt.Println("Sub-slice:", subSlice)

	// Modificando o subslice
	subSlice[0] = 99

	fmt.Println("Após modificar sub-slice:")
	fmt.Println("Original:", original)
	fmt.Println("Sub-slice:", subSlice)
}