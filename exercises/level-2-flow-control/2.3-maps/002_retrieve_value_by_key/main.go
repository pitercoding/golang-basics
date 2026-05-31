package main

import "fmt"

func main() {
	studentGrades := map[string]int{
		"Racha Cuca":     90,
		"Quase Nada":     85,
		"Poucas Trancas": 95,
	}

	student := "Poucas Trancas"
	grade, exists := studentGrades[student]

	if exists {
		fmt.Printf("%s's grade: %d\n", student, grade)
	} else {
		fmt.Println("Student not found")
	}
}

