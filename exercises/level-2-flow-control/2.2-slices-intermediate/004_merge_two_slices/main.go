package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	merged := append(slice1, slice2...)

	fmt.Println("Slice 1: ", slice1)
	fmt.Println("Slice 2: ", slice2)
	fmt.Println("Merged slice: ", merged)
}
