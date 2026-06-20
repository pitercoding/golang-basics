package main

import "fmt"

func main() {
	first := []int{1, 2, 3, 4, 5}
	second := []int{1, 2, 4, 4, 5, 6}

	result := compareSlices(first, second)
	
	fmt.Println("First:", first)
	fmt.Println("Second:", second)
	fmt.Println("Equal:", result)
}

func compareSlices(first, second []int) bool {
	if len(first) != len(second) {
		return false
	}

	for i := range first {
		if first[i] != second[i] {
			return false
		}
	}

	return true
}