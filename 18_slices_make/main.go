package main

import "fmt"

func main() {
	slice := make([]int, 0, 5) // len 0, cap 5
	fmt.Printf("Inicial: Slice: %v, Len: %d, Cap: %d\n", slice, len(slice), cap(slice))

	// Crescendo dentro da capacidade inicial
	for i := 1; i <= 8; i++ {
		slice = append(slice, i)
		fmt.Printf("After append %d: Slice: %v, Len: %d, Cap: %d\n", i, slice, len(slice), cap(slice))
	}

	// Crescendo além da capacidade inicial
	slice = append(slice, 6)
	fmt.Printf("After append 6: Slice: %v, Len: %d, Cap: %d\n", slice, len(slice), cap(slice))
}