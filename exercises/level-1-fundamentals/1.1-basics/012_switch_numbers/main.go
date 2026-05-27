package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter a number between 1 and 3: ")
	fmt.Scanln(&number)

	switch number {
	case 1:
		fmt.Println("You entered one!")
	case 2:
		fmt.Println("You entered two!")
	case 3:
		fmt.Println("You entered three!")
	default:
		fmt.Println("Invalid number!")
	}
}