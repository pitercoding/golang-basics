package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	target := 50

	fmt.Println("\n=== Linear Search ===")
	fmt.Println("Slice:", numbers)

	index, found := linearSearch(numbers, target)

	if !found {
		fmt.Printf("Target %d not found.\n", target)
	} else {
		fmt.Printf("Target %d found at index %d.\n", target, index)
	}

}

func linearSearch(numbers []int, target int) (int, bool) {
	for i, value := range numbers {
		if value == target {
			return i, true
		}
	}

	return -1, false
}
