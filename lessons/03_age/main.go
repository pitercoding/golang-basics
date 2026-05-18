package main

import "fmt"

func main() {
	var age int

	fmt.Print("Inform your age: ")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("You are over 18.")
	} else {
		fmt.Println("You are under 18.")
	}
}