package main

import "fmt"

func main() {
	numbers := []int{}

	if len(numbers) == 0 {
		fmt.Println("Slice is empty")
	} else {
		fmt.Println("Slice is NOT empty")
	}
}