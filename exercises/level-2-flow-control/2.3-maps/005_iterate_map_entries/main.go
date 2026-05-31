package main

import "fmt"

func main() {
	studentGrades := map[string]int{
		"Alice": 90,
		"Bob":   85,
		"Carol": 95,
	}

	for name, grade := range studentGrades {
		fmt.Printf("%s -> %d\n", name, grade)
	}
}
