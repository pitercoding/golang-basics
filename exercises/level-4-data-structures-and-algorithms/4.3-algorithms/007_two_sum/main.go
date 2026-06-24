package main

import "fmt"

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9

	first, second, found := twoSum(numbers, target)

	fmt.Println("\n=== Two Sum ===")
	fmt.Println("Numbers:", numbers)
	fmt.Println("Target:", target)

	if found {
		fmt.Printf("Found at indexes %d and %d\n", first, second)
		fmt.Printf("Values: %d + %d = %d\n",
			numbers[first],
			numbers[second],
			target,
		)
	} else {
		fmt.Println("No valid pair found.")
	}
}

func twoSum(numbers []int, target int) (int, int, bool) {
	seen := make(map[int]int)

	for index, value := range numbers {
		complement := target - value

		if previousIndex, exists := seen[complement]; exists {
			return previousIndex, index, true
		}

		seen[value] = index
	}

	return -1, -1, false
}
