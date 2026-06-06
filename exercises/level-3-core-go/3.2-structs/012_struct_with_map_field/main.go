package main

import "fmt"

type Student struct {
	Name   string
	Grades map[string]float64
}

func main() {
	var name string
	var totalSubjects int

	fmt.Println("\n=== Student Subjects ===")

	fmt.Print("Enter student name: ")
	fmt.Scanln(&name)

	fmt.Print("How many subjects? ")
	fmt.Scanln(&totalSubjects)

	if totalSubjects <= 0 {
		fmt.Println("Invalid number of subjects")
		return
	}

	grades := make(map[string]float64)

	for i := 0; i < totalSubjects; i++ {
		var subject string
		var grade float64

		fmt.Printf("Enter subject %d: ", i+1)
		fmt.Scanln(&subject)

		fmt.Printf("Enter grade for %s: ", subject)
		fmt.Scanln(&grade)

		if grade < 0 || grade > 10 {
			fmt.Println("Invalid grade")
			i--
			continue
		}

		grades[subject] = grade
	}

	student := Student{
		Name:   name,
		Grades: grades,
	}

	fmt.Println("\n--- Student Information ---")
	fmt.Printf("Name: %s\n", student.Name)

	for subject, grade := range student.Grades {
		fmt.Printf("%s: %.2f\n", subject, grade)
	}
}
