package main

import "fmt"

func main() {
	products := map[string]float64{
		"iPhone": 3204.33,
		"iMac":   5389.99,
		"iWatch": 300.00,
		"AirPod": 149.99,
	}

	var total float64
	totalOfProducts := len(products)

	if len(products) == 0 {
		fmt.Println("No products")
		return
	}

	for productName, price := range products {
		total += price
		fmt.Printf("%s: %.2f€\n", productName, price)
	}

	average := total / float64(totalOfProducts)

	fmt.Printf("Total: %.2f€\n", total)
	fmt.Printf("Average: %.2f€\n", average)
}
