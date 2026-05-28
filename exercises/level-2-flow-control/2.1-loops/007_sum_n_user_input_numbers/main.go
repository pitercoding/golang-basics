package main

import "fmt"

func main() {
	var totalNumbers int
	var number int
	sum := 0

	fmt.Print("How many numbers do you want to sum? ")
	fmt.Scan(&totalNumbers)

	for i := 1; i <= totalNumbers; i++ {
		fmt.Printf("Enter number %d: ", i)
		fmt.Scan(&number)

		sum += number
	}

	fmt.Printf("Total sum: %d\n", sum)
}
