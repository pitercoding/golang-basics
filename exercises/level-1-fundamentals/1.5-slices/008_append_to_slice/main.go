package main

import "fmt"

func main() {
	var numbers []int

	numbers = append(numbers, 10)
	numbers = append(numbers, 20)
	numbers = append(numbers, 30)
	numbers = append(numbers, 40)

	fmt.Println("Slice:", numbers)
}