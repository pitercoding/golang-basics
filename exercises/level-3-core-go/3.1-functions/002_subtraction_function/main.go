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

	fmt.Println("\n=== Subtraction Function ===")

	n1 := readFloat(scanner, "Enter first number: ")
	n2 := readFloat(scanner, "Enter second number: ")

	result := subtract(n1, n2)

	fmt.Printf("%.2f - %.2f = %.2f\n", n1, n2, result)
}

func subtract(a, b float64) float64 {
	return a - b
}

func readFloat(scanner *bufio.Scanner, message string) float64 {
	for {
		fmt.Print(message)

		if !scanner.Scan() {
			fmt.Println("Error reading input")
			return 0
		}

		input := strings.TrimSpace(scanner.Text())

		value, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return value
		}

		fmt.Println("Invalid input. Try again.")
	}
}
