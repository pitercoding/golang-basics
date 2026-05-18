package main

import "fmt"

func main() {
	array := [5]int{10, 20, 30, 40, 50}

	slice := array[1:4]
	fmt.Printf("Array: %v\n", array)
	fmt.Printf("Slice: %v, Len: %d, Cap: %d\n", slice, len(slice), cap(slice))

	// Modificando o slice afeta o array subjacente
	slice[0] = 100
	fmt.Printf("Após modificar o slice:\n")
	fmt.Printf("Array: %v\n", array)
	fmt.Printf("Slice: %v\n", slice)
}