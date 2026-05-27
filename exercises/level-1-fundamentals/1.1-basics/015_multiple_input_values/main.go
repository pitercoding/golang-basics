package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Enter your name and age: ")
	fmt.Scanln(&name, &age)

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
}