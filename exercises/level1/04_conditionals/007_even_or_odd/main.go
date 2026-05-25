package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	checkEvenOrOdd(number)
}

func checkEvenOrOdd(n int) {
	if n % 2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}
}