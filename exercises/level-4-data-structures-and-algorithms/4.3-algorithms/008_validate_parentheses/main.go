package main

import "fmt"

func main() {
	examples := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
	}

	fmt.Println("\n=== Validate Parentheses ===")

	for _, example := range examples {
		fmt.Printf("%s -> %t\n",
			example,
			isValid(example),
		)
	}
}

func isValid(text string) bool {
	stack := []rune{}

	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range text {

		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		top := stack[len(stack)-1]

		if top != pairs[char] {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
