package main

import "fmt"

type Student struct {
	Name   string
	Grades []float64
}

func (s Student) Average() float64 {
	if len(s.Grades) == 0 {
		return 0
	}

	sum := 0.0

	for _, grade := range s.Grades {
		sum += grade
	}

	return sum / float64(len(s.Grades))
}

func main() {
	var name string
	var totalGrades int

	fmt.Println("\n=== Student Grades ===")

	fmt.Print("Enter student name: ")
	fmt.Scanln(&name)

	fmt.Print("How many grades? ")
	fmt.Scanln(&totalGrades)

	if totalGrades <= 0 {
		fmt.Println("Invalid number of grades")
		return
	}

	grades := make([]float64, 0, totalGrades)

	for i := 0; i < totalGrades; i++ {
		var grade float64

		fmt.Printf("Enter grade %d: ", i+1)
		fmt.Scanln(&grade)

		if grade < 0 || grade > 10 {
			fmt.Println("Invalid grade")
			i--
			continue
		}

		grades = append(grades, grade)
	}

	student := Student{
		Name:   name,
		Grades: grades,
	}

	fmt.Println("\n--- Student Information ---")
	fmt.Printf("Name: %s\n", student.Name)
	fmt.Printf("Grades: %v\n", student.Grades)
	fmt.Printf("Average: %.2f\n", student.Average())
}
