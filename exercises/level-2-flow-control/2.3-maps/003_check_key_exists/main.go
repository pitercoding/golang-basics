package main

import "fmt"

func main() {
	studentsGrades := map[string]int{
		"Alice": 80,
		"Bob": 70,
		"Lia": 98,
	}

	student := "Jhon"

	grade, exists := studentsGrades[student]

	if exists {
		fmt.Printf("%s's grade: %d\n", student, grade)
	} else {
		fmt.Println("Student not found!")
	}
}
