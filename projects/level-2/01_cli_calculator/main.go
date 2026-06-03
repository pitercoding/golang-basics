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
		fmt.Println("\n=== CLI Calculator ===")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("0. Exit")

		option := readInt(scanner, "Choose an option: ")

		switch option {
		case 1:
			n1 := readFloat(scanner, "Enter first number: ")
			n2 := readFloat(scanner, "Enter second number: ")

			fmt.Printf("Result: %.2f\n", add(n1, n2))

		case 2:
			n1 := readFloat(scanner, "Enter first number: ")
			n2 := readFloat(scanner, "Enter second number: ")

			fmt.Printf("Result: %.2f\n", subtract(n1, n2))

		case 3:
			n1 := readFloat(scanner, "Enter first number: ")
			n2 := readFloat(scanner, "Enter second number: ")

			fmt.Printf("Result: %.2f\n", multiply(n1, n2))

		case 4:
			n1 := readFloat(scanner, "Enter first number: ")
			n2 := readFloat(scanner, "Enter second number: ")

			result, err := divide(n1, n2)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Printf("Result: %.2f\n", result)

		case 0:
			fmt.Println("\nThank you for using the calculator. Bye!")
			return

		default:
			fmt.Println("\nInvalid option! Try again.")
		}
	}
}

func add(n1, n2 float64) float64 {
	return n1 + n2
}

func subtract(n1, n2 float64) float64 {
	return n1 - n2
}

func multiply(n1, n2 float64) float64 {
	return n1 * n2
}

func divide(n1, n2 float64) (float64, error) {
	if n2 == 0 {
		return 0, fmt.Errorf("[ERROR] Cannot divide by zero!")
	}

	return n1 / n2, nil
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

func readFloat(scanner *bufio.Scanner, message string) float64 {
	for {
		fmt.Print(message)

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		value, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return value
		}

		fmt.Println("Invalid input. Try again.")
	}
}