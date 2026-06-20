package main

import "fmt"

func main() {
	grades := make(map[string]map[string]int)

	grades["Jhon"] = make(map[string]int)
	grades["Mary"] = make(map[string]int)

	grades["Jhon"]["Math"] = 90
	grades["Jhon"]["English"] = 85

	grades["Mary"]["Math"] = 95
	grades["Mary"]["English"] = 88

	for student, subjects := range grades {
		fmt.Println(student)

		for subject, grade := range subjects {
			fmt.Printf("  %s: %d\n", subject, grade)
		}
	}
}
