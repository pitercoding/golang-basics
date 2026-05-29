package main

import "fmt"

func main() {
	numbers := []int{1, 2, 2, 3, 4, 4, 5}
	seen := make(map[int]bool)
	result := []int{}

	for _, n := range numbers {
		if !seen[n] {
			seen[n] = true
			result = append(result, n)
		}
	}

	fmt.Println("Without duplicates:", result)
}
