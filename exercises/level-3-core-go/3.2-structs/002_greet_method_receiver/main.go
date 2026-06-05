package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func (p Person) isAdult() bool {
	return p.Age >= 18
}

func main() {
	person := Person{
		Name: "Chapolin",
		Age: 29,
	}

	person.Greet()

	if person.isAdult() {
		fmt.Println("Adult!")
	}
}