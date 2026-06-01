package main

import "fmt"

func main() {
	items := []string{"go", "go", "rust", "go", "java", "typescript"}

	counter := map[string]int{}

	for _, item := range items {
		counter[item]++
	}

	for key, occurrence := range counter {
		fmt.Printf("%s -> %d\n", key, occurrence)
	}
}
