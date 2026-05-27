package main

import (
	"fmt"
	"slices"
)

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}

	newSlice := slices.Concat(slice1, slice2)

	fmt.Println("Concatenated slice:", newSlice)

}