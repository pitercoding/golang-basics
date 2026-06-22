package main

import (
	"fmt"
)

func main() {
	sales := map[string]float64{
		"Laptop":   1200,
		"Mouse":    50,
		"Keyboard": 80,
	}

	result := aggregateSales(sales)

	fmt.Println("\n=== Sales ($) ===")

	for product, value := range sales {
		fmt.Printf("%s: %.2f\n", product, value)
	}

	fmt.Printf("Total: %.2f\n", result)
	fmt.Printf("Average: %.2f\n", averageSales(sales))
	fmt.Printf("Highest: %.2f\n", highestSale(sales))
	fmt.Printf("Lowest: %.2f\n", lowestSale(sales))
}

func aggregateSales(sales map[string]float64) float64 {
	total := 0.0

	for _, value := range sales {
		total += value
	}

	return total
}

func averageSales(sales map[string]float64) float64 {
	total := 0.0
	count := 0

	for _, value := range sales {
		total += value
		count++
	}

	average := total / float64(count)

	return average
}

func highestSale(sales map[string]float64) float64 {
	if len(sales) == 0 {
		return 0
	}

	var max float64
	first := true

	for _, value := range sales {
		if first || value > max {
			max = value
			first = false
		}
	}

	return max
}

func lowestSale(sales map[string]float64) float64 {
	if len(sales) == 0 {
		return 0
	}

	var min float64
	first := true

	for _, value := range sales {
		if first || value < min {
			min = value
			first = false
		}
	}

	return min
}
