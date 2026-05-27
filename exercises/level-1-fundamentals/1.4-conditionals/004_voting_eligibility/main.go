package main

import "fmt"

func main() {
	var age int

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	checkVotingEligibility(age)
}

func checkVotingEligibility(age int) {
	if age >= 16 {
		fmt.Println("Eligible to vote")
	} else {
		fmt.Println("Not eligible to vote")
	}
}