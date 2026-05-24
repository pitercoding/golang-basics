package main

import "fmt"

func main() {
	var number float64

	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	checkNumber(number)
}

func checkNumber(number float64) {
	if number > 0 {
		fmt.Println("Positive number")
	} else if number < 0 {
		fmt.Println("Negative number")
	} else {
		fmt.Println("Zero")
	}
}