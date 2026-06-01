package main

import "fmt"

func main() {
	items := map[string]string{
		"Apple":      "fruits",
		"Banana":     "fruits",
		"BMW":        "cars",
		"Rio":        "cities",
		"Tesla":      "cars",
		"Watermelon": "fruits",
		"Berlin":     "cities",
	}
	
	data := map[string][]string{}

	for item, category := range items {
		data[category] = append(data[category], item)
	}

	for category, list := range data {
		fmt.Println(category, "->", list)
	}
}
