package main

import "fmt"

func main() {
	var slice []int

	fmt.Printf("Before append: Slice: %v, Len: %d, Cap: %d\n", slice, len(slice), cap(slice))

	for i := 1; i <= 6; i++ {
		slice = append(slice, i)
		fmt.Printf("After append %d: Slice: %v, Len: %d, Cap: %d\n", i, slice, len(slice), cap(slice))
	}
}