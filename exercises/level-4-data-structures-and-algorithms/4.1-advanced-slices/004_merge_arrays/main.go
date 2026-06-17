package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=== Merge Arrays ===")

		quantity1 := readInt(scanner, "How many numbers in first array? (0 to exit): ")

		if quantity1 == 0 {
			fmt.Println("\nGoodbye!")
			return
		}

		if quantity1 < 0 {
			fmt.Println("[ERROR] Only positive numbers are accepted! Try again.")
			continue
		}

		numbers1 := make([]int, 0, quantity1)

		for i := 0; i < quantity1; i++ {
			number1 := readInt(
				scanner,
				fmt.Sprintf("Enter number %d: ", i+1),
			)
			numbers1 = append(numbers1, number1)
		}

		quantity2 := readInt(scanner, "How many numbers in second array? (0 to exit): ")

		if quantity2 == 0 {
			fmt.Println("\nGoodbye!")
			return
		}

		if quantity2 < 0 {
			fmt.Println("[ERROR] Only positive numbers are accepted! Try again.")
			continue
		}

		numbers2 := make([]int, 0, quantity2)

		for i := 0; i < quantity2; i++ {
			number2 := readInt(
				scanner,
				fmt.Sprintf("Enter number %d: ", i+1),
			)
			numbers2 = append(numbers2, number2)
		}

		merged := mergeArrays(numbers1, numbers2)

		fmt.Printf("\nFirst array:  %v\n", numbers1)
		fmt.Printf("Second array: %v\n", numbers2)
		fmt.Printf("Merged array: %v\n", merged)
	}
}

func mergeArrays(
	first []int,
	second []int,
) []int {
	merged := make([]int, 0, len(first)+len(second))

	merged = append(merged, first...)
	merged = append(merged, second...)

	return merged
}

func readInt(scanner *bufio.Scanner, message string) int {
	for {
		fmt.Print(message)

		scanner.Scan()

		input := strings.TrimSpace(scanner.Text())

		value, err := strconv.Atoi(input)
		if err == nil {
			return value
		}

		fmt.Println("Invalid value. Try again!")
	}
}
