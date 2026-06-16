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
		fmt.Println("\n=== Second Largest Element ===")

		quantity := readInt(scanner, "How many numbers? (Enter 0 to exit): ")

		if quantity == 0 {
			fmt.Println("\nGoodbye!")
			return
		}

		if quantity < 2 {
			fmt.Println("You must enter at least 2 numbers.")
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

		fmt.Printf("\nNumbers: %v\n", numbers)

		secondLargest, found := findSecondLargest(numbers)

		if found {
			fmt.Printf("Second largest: %d\n", secondLargest)
		} else {
			fmt.Println("Second largest number not found.")
		}

	}

}

func findSecondLargest(
	numbers []int,
) (int, bool) {

	if len(numbers) < 2 {
		return -1, false
	}

	largest := numbers[0]
	secondLargest := numbers[1]

	if secondLargest > largest {
		largest, secondLargest = secondLargest, largest
	}

	for i := 2; i < len(numbers); i++ {
		if numbers[i] > largest {
			secondLargest = largest
			largest = numbers[i]
		} else if numbers[i] > secondLargest {
			secondLargest = numbers[i]
		}
	}

	return secondLargest, true
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

		fmt.Print("Invalid number. Try again.\n\n")
	}
}
