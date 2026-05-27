package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "orange", "mango"}

	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
}