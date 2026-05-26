package main

import "fmt"

func main() {
	var password string

	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	validatePassword(password)
}

func validatePassword(password string)  {
	if len(password) >= 8 {
		fmt.Println("Valid password")
	} else {
		fmt.Println("Password must contain at least 8 characters")
	}
}