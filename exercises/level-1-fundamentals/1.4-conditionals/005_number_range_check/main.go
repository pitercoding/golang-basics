package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	checkRange(n)
}

func checkRange(n int) {
	if n >= 10 && n <= 20 {
		fmt.Println("Number is between 10 and 20")
	} else {
		fmt.Println("Number is outside the range")
	}
}