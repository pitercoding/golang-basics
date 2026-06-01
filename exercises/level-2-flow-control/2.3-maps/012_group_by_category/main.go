package main

import "fmt"

func main() {
	items := map[string]string{
		"Apple":      "fruit",
		"Banana":     "fruit",
		"Carrot":     "vegetable",
		"Potato":     "vegetable",
		"Watermelon": "fruit",
		"Lettuce":    "vegetable",
	}

	grouped := map[string][]string{}

	for item, category := range items {
		grouped[category] = append(grouped[category], item)
	}

	for category, list := range grouped {
		fmt.Println(category, "→", list)
	}
}
