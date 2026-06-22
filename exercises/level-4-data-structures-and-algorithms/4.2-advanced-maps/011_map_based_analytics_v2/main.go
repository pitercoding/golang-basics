package main

import (
	"fmt"
)

type Analytics struct {
	Count    int
	Sum      float64
	Avg      float64
	Max      float64
	Min      float64
	AboveAvg map[string]float64
}

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

	analytics := generateAnalytics(sales)

	fmt.Println("\n=== Analytics ===")
	fmt.Printf("Product Quantity: %d\n", analytics.Count)
	fmt.Printf("Sum($): %.2f\n", analytics.Sum)
	fmt.Printf("Average($): %.2f\n", analytics.Avg)
	fmt.Printf("Highest Sale($): %.2f\n", analytics.Max)
	fmt.Printf("Lowest Sale($): %.2f\n", analytics.Min)

	fmt.Println("\n=== Sales Above Average ===")
	for product, price := range analytics.AboveAvg {
		fmt.Printf("%s: %.2f\n", product, price)
	}
}

func generateAnalytics(sales map[string]float64) Analytics {
	a := Analytics{
		Count:    len(sales),
		AboveAvg: make(map[string]float64),
	}

	if len(sales) == 0 {
		return a
	}

	first := true

	for _, price := range sales {
		a.Sum += price

		if first {
			a.Max = price
			a.Min = price
			first = false
		} else {
			if price > a.Max {
				a.Max = price
			}
			if price < a.Min {
				a.Min = price
			}
		}
	}

	a.Avg = a.Sum / float64(a.Count)

	for product, price := range sales {
		if price > a.Avg {
			a.AboveAvg[product] = price
		}
	}

	return a
}
