package main

import "fmt"

func main() {
	var age int

	fmt.Println("\n=== Age Group Classifier ===")
	fmt.Print("Enter age: ")
	fmt.Scanln(&age)

	if age < 0 || age > 130 {
		fmt.Println("Age must be between 0 and 130!")
		return
	}

	group := classifyAgeGroup(age)
	fmt.Printf("\nAge group: %s\n", group)
}

func classifyAgeGroup(age int) string {

	if age <= 2 {
		return "Baby"
	}
	
	if age <= 12 {
		return "Child"
	}

	if age <= 17 {
		return "Teenager"
	}

	if age <= 39 {
		return "Young Adult"
	}

	if age <= 59 {
		return "Adult"
	}

	return "Senior"
}