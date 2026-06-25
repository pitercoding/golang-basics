package main

import "fmt"

func main() {
	tests := [][2]string{
		{"listen", "silent"},
		{"hello", "world"},
		{"heart", "earth"},
		{"angel", "glean"},
	}

	fmt.Println("\n=== Check Anagram Strings ===")

	for _, test := range tests {
		result := isAnagram(test[0], test[1])

		fmt.Printf(
			"%s | %s -> %t\n",
			test[0],
			test[1],
			result,
		)
	}

}

func isAnagram(first, second string) bool {
	if len(first) != len(second) {
		return false
	}

	counts := make(map[rune]int)

	for _, char := range first {
		counts[char]++
	}

	for _, char := range second {
		counts[char]--
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}
