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

	var grades []float64
	var totalGrades int

	fmt.Println("\n=== Grade Calculator ===")
	fmt.Print("How many grades do you want to enter? ")

	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	value, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	totalGrades = value

	if totalGrades <= 0 {
		fmt.Println("Invalid number of grades")
		return
	}

	for i := 0; i < totalGrades; i++ {
		for {
			fmt.Printf("Enter grade %d: ", i+1)

			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())

			grade, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Invalid input, try again")
				continue
			}

			if grade < 0 || grade > 10 {
				fmt.Println("Invalid grade, try again")
				continue
			}

			grades = append(grades, grade)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	average := calculateAverage(grades)
	status := classifyGrade(average)

	fmt.Println()
	fmt.Printf("Average: %.2f\n", average)
	fmt.Printf("Status: %s\n", status)
}

func calculateAverage(grades []float64) float64 {
	if len(grades) == 0 {
		return 0
	}

	var sum float64

	for _, grade := range grades {
		sum += grade
	}

	return sum / float64(len(grades))
}

func classifyGrade(score float64) string {
	if score >= 7 {
		return "Approved"
	}

	if score >= 5 {
		return "Recovery"
	}

	return "Failed"
}
