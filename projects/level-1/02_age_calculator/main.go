package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your birth year: ")
	scanner.Scan()

	input := strings.TrimSpace(scanner.Text())

	birthYear, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid year!")
		return
	}

	currentYear := time.Now().Year()

	if birthYear > currentYear {
		fmt.Println("Birth year cannot be in the future")
		return
	}

	age := calculateAge(birthYear)

	if age < 0 || age > 130 {
		fmt.Println("Invalid age")
		return
	}

	fmt.Printf("You are %d years old\n", age)
}

func calculateAge(birthYear int) int {
	currentYear := time.Now().Year()
	return currentYear - birthYear
}
