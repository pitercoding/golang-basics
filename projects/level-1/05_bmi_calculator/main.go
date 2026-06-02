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

	fmt.Println("\n=== BMI Calculator ===")

	weight := readFloat(scanner, "Enter your weight (kg): ")
	height := readFloat(scanner, "Enter your height (m): ")

	if height <= 0 {
		fmt.Println("Invalid height!")
		return
	}

	bmi := calculateBMI(weight, height)
	status := classifyBMI(bmi)

	fmt.Println()
	fmt.Printf("BMI: %.2f\n", bmi)
	fmt.Printf("Status: %s\n", status)
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

		fmt.Println("Invalid input, try again")
	}
}

func calculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}

func classifyBMI(bmi float64) string {
	switch {
	case bmi < 18.5:
		return "Underweight"
	case bmi < 25:
		return "Normal weight"
	case bmi < 30:
		return "Overweight"
	default:
		return "Obesity"
	}
}
