package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	checkNumber(number)
}

func checkNumber(number int) {
	if number > 0 && number%2 == 0 {
		fmt.Println("Positive and even number")
	} else {
		fmt.Println("Does not meet the condition")
	}
}
