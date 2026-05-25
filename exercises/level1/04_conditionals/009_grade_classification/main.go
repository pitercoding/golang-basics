package main

import "fmt"

func main() {
	var score int

	fmt.Print("Enter score (0-100): ")
	fmt.Scanln(&score)

	classifyGrade(score)
}

func classifyGrade(score int) {
	if score < 0 || score > 100 {
		fmt.Println("Invalid score!")
		return
	}


	if score >= 90 && score <= 100 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}
}
