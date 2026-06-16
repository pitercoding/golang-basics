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

	numbers := []int{
		10,
		20,
		30,
		40,
		50,
		60,
		70,
		80,
		90,
	}

	for {
		fmt.Println("\n=== Binary Search ===")
		fmt.Printf("Numbers : %v\n", numbers)

		target := readInt(scanner, "Enter target number (0 to exit): ")

		if target < 0 {
			fmt.Println("[ERROR] Target must be greater than 0.")
			continue
		}

		if target == 0 {
			fmt.Println("\nLeaving program. Thank you...")
			return
		}

		index, found := binarySearch(numbers, target)

		if found {
			fmt.Printf("Found at index %d\n", index)
		} else {
			fmt.Println("Number not found.")
		}
	}

}

func binarySearch(
	numbers []int,
	target int,
) (int, bool) {

	left := 0
	right := len(numbers) - 1

	for left <= right {

		mid := (left + right) / 2

		if numbers[mid] == target {
			return mid, true
		}

		if target > numbers[mid] {
			left = mid + 1
		} else {
			right = mid - 1
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

		fmt.Print("Invalid number. Try again.\n\n")
	}
}
