package main

import "fmt"

func main() {
	numbers := []int{-5, 10, -3, 8, 0, 15, -1}

	positiveNumbers := []int{}

	for _, number := range numbers {
		if number > 0 {
			positiveNumbers = append(positiveNumbers, number)
		}
	}

	fmt.Println("Positive numbers:", positiveNumbers)
}
