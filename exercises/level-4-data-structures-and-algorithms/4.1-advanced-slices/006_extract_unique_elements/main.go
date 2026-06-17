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
		fmt.Println("\n=== Extract Unique Elements ===")

		quantity := readInt(scanner, "How many numbers? (0 to exit): ")

		if quantity == 0 {
			fmt.Println("\nGoodbye!")
			return
		}

		if quantity < 0 {
			fmt.Println("[ERROR] Only positive numbers are accepted! Try again.")
			continue
		}

		numbers := make([]int, 0, quantity)

		for i := 0; i < quantity; i++ {
			number1 := readInt(
				scanner,
				fmt.Sprintf("Enter number %d: ", i+1),
			)

			numbers = append(numbers, number1)
		}

		result := extractUnique(numbers)

		fmt.Println("Numbers: ", numbers)
		fmt.Println("Unique Elements:", result)
	}
}

func extractUnique(numbers []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, number := range numbers {
		if !seen[number] {
			result = append(result, number)
			seen[number] = true
		}
	}

	return result
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

		fmt.Println("Invalid value! Try agin.")
	}
}
