package main

import (
	"fmt"
	"time"
)

func main() {
	var age int
	for {
		fmt.Print("Enter your age: ")
		_, err := fmt.Scanln(&age)

		if err != nil {
			fmt.Println("Invalid input! Only numbers are allowed.")
			var discard string
			fmt.Scanln(&discard) // clear remaining input
			continue
		}

		if age < 0 || age > 120 {
			fmt.Println("Invalid age! It must be between 0 and 120.")
			continue
		}
		break
	}

	birthYear := calculateBirthYear(age)

	fmt.Printf("You are %d years old and were born in %d\n", age, birthYear)
}

func calculateBirthYear(age int) int {
	currentYear := time.Now().Year()
	return currentYear - age
}