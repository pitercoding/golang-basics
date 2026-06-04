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

	fmt.Println("\n=== Sum Two Numbers ===")

	n1 := readInt(scanner, "Enter first number: ")
	n2 := readInt(scanner, "Enter second number: ")

	result := sum(n1, n2)

	fmt.Printf("%d + %d = %d\n", n1, n2, result)
}

func sum(a, b int) int {
	return a + b
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

		fmt.Println("Invalid input. Try again.")
	}

}
