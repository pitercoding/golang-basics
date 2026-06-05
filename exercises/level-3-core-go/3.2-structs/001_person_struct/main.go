package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var name string
	var age int

	fmt.Println("\n=== Person Struct ===")

	fmt.Print("Enter name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter age: ")
	fmt.Scanln(&age)

	person := Person{
		Name: name,
		Age:  age,
	}

	fmt.Println("\nPerson Information")
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
}
