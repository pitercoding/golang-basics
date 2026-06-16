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
		fmt.Println("\n=== Linear Search ===")
		quantity := readInt(scanner, "How many numbers (0 to exit)? ")

		if quantity < 0 {
			fmt.Println("Quantity must be greater than zero.")
			continue
		}

		if quantity == 0 {
			fmt.Println("\nLeaving program. Thank you...")
			return
		}

		numbers := make([]int, 0, quantity)

		for i := 0; i < quantity; i++ {
			number := readInt(
				scanner, 
				fmt.Sprintf("Enter number %d: ", i+1),
			)

			numbers = append(numbers, number)
		}

		fmt.Printf("\nNumbers: %v\n", numbers)
		target := readInt(scanner, "Enter target number: ")

		index, found := linearSearch(numbers, target)

		if found {
			fmt.Printf("Found at index %d\n", index)
		} else {
			fmt.Println("Number not found.")
		}
	}
}

func linearSearch(numbers []int, target int) (int, bool) {
	for index, value := range numbers {
		if value == target {
			return index, true
		}
	}

	return -1, false
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

		fmt.Println("\nInvalid number. Try again.")
	}
}