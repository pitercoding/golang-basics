package main

import "fmt"

func main() {
	scores := map[string]int{
		"John":  90,
		"Mary":  65,
		"Peter": 80,
		"Anna":  55,
	}

	result := filterScores(scores, 70)

	fmt.Println("Original:", scores)
	fmt.Println("Filtered:", result)
}

func filterScores(
	scores map[string]int,
	minScore int,
) map[string]int {

	result := make(map[string]int)

	for name, score := range scores {
		if score >= minScore {
			result[name] = score
		}
	}

	return result
}
