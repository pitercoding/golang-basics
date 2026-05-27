package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	evenCount := 0
	oddCount := 0

	var evenNumbers []int
	var oddNumbers []int

	for _, n := range numbers {
		if n%2 == 0 {
			evenCount++
			evenNumbers = append(evenNumbers, n)
		} else {
			oddCount++
			oddNumbers = append(oddNumbers, n)
		}
	}

	fmt.Printf("Even numbers: count = %d, final slice = %v\n", evenCount, evenNumbers)
	fmt.Printf("Odd numbers: count = %d, final slice = %v\n", oddCount, oddNumbers)
}