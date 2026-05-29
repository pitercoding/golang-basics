package main

import "fmt"

func main() {
	var count int
	var number int
	sum := 0

	fmt.Print("How many numbers? ")
	fmt.Scan(&count)

	for i := 1; i <= count; i++ {
		fmt.Printf("Enter number %d: ", i)
		fmt.Scan(&number)

		sum += number
	}

	average := float64(sum) / float64(count)

	fmt.Printf("Average: %.2f\n", average)
}
