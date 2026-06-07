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

func (s Student) Status() string {
	average := s.Average()

	if average >= 7.0 {
		return "Approved"
	}

	if average >= 5.0 {
		return "Recovery"
	}

	return "Failed"
}

func main() {
	student := Student {
		Name: "Bia",
		Grades: []float64 {
			10,
			8,
			6,
		},
	}

	fmt.Println("=== School System ===")
	fmt.Printf("Name: %s\n", student.Name)
	fmt.Printf("Average: %.2f\n", student.Average())
	fmt.Printf("Status: %s\n", student.Status())
}