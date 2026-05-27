package main

import (
	"fmt"
)

func main() {
	numbers := []int{10, 20, 30, 40, 50}

	var value int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&value)

	fmt.Printf("Searching for %d...\n", value)

	if contains(numbers, value) {
		fmt.Println("Value found!")
	} else {
		fmt.Println("Value not found!")
	}
}

func contains(slice []int, target int) bool {
	for _, n := range slice {
		if n == target {
			return true
		}
	}
	return false
}
