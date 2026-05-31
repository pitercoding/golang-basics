package main

import "fmt"

func main() {
	studentGrades := map[string]int{
		"Alice": 90,
		"Bob":   85,
		"Carol": 95,
	}

	fmt.Println("Before deletion:")
	fmt.Println(studentGrades)

	delete(studentGrades, "Bob")

	if _, exists := studentGrades["Bob"]; !exists {
		fmt.Println("Bob removed successfully")
	}

	fmt.Println("After deletion:")
	fmt.Println(studentGrades)
}
