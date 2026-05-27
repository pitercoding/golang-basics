package main

import "fmt"

func main() {
	numbers := []float64{10, 20, 30, 40, 50}

	sum := 0.0

	for _, n := range numbers {
		sum += n
	}

	average := sum / float64(len(numbers))

	fmt.Printf("Average: %.2f\n", average)
}