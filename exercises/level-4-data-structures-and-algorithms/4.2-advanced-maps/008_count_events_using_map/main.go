package main

import "fmt"

func main() {
	events := []string{
		"login",
		"logout",
		"login",
		"purchase",
		"login",
		"logout",
	}

	counts := countEvents(events)
	fmt.Println(counts)
}

func countEvents(events []string) map[string]int {
	counts := make(map[string]int)

	for _, event := range events {
		counts[event]++
	}

	return counts
}