package main

import "fmt"

type entry struct {
	product string
	value   int
}

func main() {
	data := []entry{
		{"iPhone", 3000},
		{"iPhone", 200},
		{"iWatch", 500},
		{"iMac", 5000},
		{"iWatch", 100},
		{"iMac", 1000},
	}

	result := map[string]int{}

	for _, price := range data {
		result[price.product] += price.value
	}

	for product, total := range result {
		fmt.Printf("%s -> %d\n", product, total)
	}
}
