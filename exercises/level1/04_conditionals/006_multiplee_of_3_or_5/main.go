package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	checkMultiple(number)
}

func checkMultiple(number int)  {
	if number % 3 || number % 3 == 0 {
		fmt.Println("Number is multiple of 3 or 5")
	} else {
		fmt.Println("Number is not multiple of 3 or 5")
	}
}