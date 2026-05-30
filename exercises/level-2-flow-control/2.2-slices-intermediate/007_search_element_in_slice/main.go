package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	target := 30

	found := false
	foundIndex := -1

	for index, number := range numbers {
		if number == target {
			found = true
			foundIndex = index
			break
		}
	}

	if found {
		fmt.Printf("Element found at index %d!\n", foundIndex)
	} else {
		fmt.Println("Element not found!")
	}
}
