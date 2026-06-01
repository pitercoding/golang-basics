package main

import "fmt"

func main() {
	products := map[string]float64{
		"iPhone": 3000,
		"iMac":   5000,
		"iWatch": 1000,
	}
	
	fmt.Println("Original Price:")
	for product, price := range products {
		fmt.Printf("%s -> %.2f\n", product, price)
	}

	fmt.Println()

	for product, price := range products {
		products[product] = price * 1.10
	}

	fmt.Println("Updated Price:")
	for product, price := range products {
		fmt.Printf("%s -> %.2f\n", product, price)
	}

}
