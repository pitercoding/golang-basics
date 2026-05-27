package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// acessar elementos
	fmt.Println("First row:", matrix[0])
	fmt.Println("Element [1][2]:", matrix[1][2])
}