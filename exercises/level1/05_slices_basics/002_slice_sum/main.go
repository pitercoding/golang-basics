package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 100}

	sum := 0

	for _, n := range numbers {
		sum += n
	}

	fmt.Println("Sum:", sum)
}