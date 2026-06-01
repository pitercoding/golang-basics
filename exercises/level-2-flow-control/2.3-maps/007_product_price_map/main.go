package main

import "fmt"

func main() {
	products := map[string]float64{
		"iPhone": 3204.33,
		"iMac":   5389.99,
		"iWatch": 300.00,
		"AirPod": 149.99,
	}

	total := 0.0

	for productName, price := range products {
		total += price
		fmt.Printf("%s: %.2f€\n", productName, price)
	}

	fmt.Printf("Total: %.2f€\n", total)
}
