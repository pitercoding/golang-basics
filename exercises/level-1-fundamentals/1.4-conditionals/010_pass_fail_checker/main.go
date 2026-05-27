package main

import "fmt"

func main() {
	var grade float64

	fmt.Print("Enter your grade: ")
	fmt.Scanln(&grade)

	checkResult(grade)
}

func checkResult(grade float64) {
	if grade < 0 || grade > 10 {
		fmt.Println("Invalid grade!")
		return
	}

	if grade >= 7 {
		fmt.Println("Approved")
	} else {
		fmt.Println("Failed")
	}
}
