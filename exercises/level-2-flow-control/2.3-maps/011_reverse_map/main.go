package main

import "fmt"

func main() {
	ages := map[string]int{
		"Alice": 20,
		"Bob":   30,
		"Carol": 40,
	}

	fmt.Println("Original:")

	for name, age := range ages {
		fmt.Printf("%s -> %d\n", name, age)
	}

	fmt.Println("")

	reversed := map[int]string{}

	for name, age := range ages {
		reversed[age] = name
	}

	fmt.Println("Reversed:")

	for age, name := range reversed {
		fmt.Printf("%d -> %s\n", age, name)
	}
}
