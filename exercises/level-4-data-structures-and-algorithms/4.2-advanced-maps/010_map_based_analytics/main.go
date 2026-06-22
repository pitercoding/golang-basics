package main

import (
	"fmt"
)

func main() {
	sales := map[string]float64{
		"Laptop":   1200,
		"Mouse":    50,
		"Keyboard": 80,
		"Monitor":  300,
		"Headset":  150,
	}

	fmt.Println("\n=== Sales ===")
	for product, price := range sales {
		fmt.Printf("%s: %.2f\n", product, price)
	}

	count, sum, avg, max, min, aboveAvg := generateAnalytics(sales)

	fmt.Println("\n=== Analytics ===")
	fmt.Printf("Product Quantity: %d\n", count)
	fmt.Printf("Sum($): %.2f\n", sum)
	fmt.Printf("Average($): %.2f\n", avg)
	fmt.Printf("Highest Sale($): %.2f\n", max)
	fmt.Printf("Lowest Sale($): %.2f\n", min)

	fmt.Println("\n=== Sales Above Average ===")
	for product, price := range aboveAvg {
		fmt.Printf("%s: %.2f\n", product, price)
	}

}

func generateAnalytics(sales map[string]float64) (count int, sum, avg, max, min float64, aboveAvg map[string]float64) {
	count = len(sales)
	sum = 0.0
	first := true

	for _, price := range sales {
		sum += price

		if first {
			max = price
			min = price
			first = false
		}

		if price > max {
			max = price
		}

		if price < min {
			min = price
		}
	}

	if count > 0 {
		avg = sum / float64(count)
	}

	aboveAvg = make(map[string]float64)

	for product, price := range sales {
		if price > avg {
			aboveAvg[product] = price
		}
	}

	return count, sum, avg, max, min, aboveAvg
}
