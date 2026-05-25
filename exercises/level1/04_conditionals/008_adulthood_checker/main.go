package main

import "fmt"

func main() {
	var age int

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	checkAdulthood(age)
}

func checkAdulthood(age int) {
	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}
}
