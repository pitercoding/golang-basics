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
		fmt.Println("\n=== Sliding Window ===")

		windowSize := readInt(scanner, "Enter Slide Window Size (0 to exit): ")

		if windowSize == 0 {
			fmt.Println("\nGoodbye!")
			return
		}

		if windowSize < 0 {
			fmt.Println("[ERROR] Only positive numbers are accepted. Try again!")
			continue
		}

		quantity := readInt(scanner, "How many numbers (0 to exit): ")

		if quantity == 0 {
			fmt.Println("\nGoodbye!")
			return
		}

		if quantity < 0 {
			fmt.Println("[ERROR] Only positive numbers are accepted. Try again!")
			continue
		}

		numbers := make([]int, 0, quantity)

		for i := 0; i < quantity; i++ {
			number := readInt(
				scanner,
				fmt.Sprintf("Enter number %d: ", i+1),
			)

			numbers = append(numbers, number)
		}

		if windowSize > len(numbers) {
			fmt.Println("Window size cannot be greater than number count.")
			continue
		}

		result := slidingWindowMaximum(numbers, windowSize)

		fmt.Println()
		fmt.Println("Numbers:", numbers)
		fmt.Println("Result:", result)
	}
}

func slidingWindowMaximum(numbers []int, windowSize int) []int {
	if windowSize <= 0 || windowSize > len(numbers) {
		return []int{}
	}

	result := []int{}

	for start := 0; start <= len(numbers)-windowSize; start++ {
		window := numbers[start : start+windowSize]

		max := findMax(window)

		result = append(result, max)
	}

	return result
}

func findMax(numbers []int) int {
	max := numbers[0]

	for _, value := range numbers[1:] {
		if value > max {
			max = value
		}
	}
	return max
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

		fmt.Println("Invalid value! Try again.")
	}
}
