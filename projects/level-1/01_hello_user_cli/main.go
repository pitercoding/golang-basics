package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var name string

	for {
		fmt.Print("Enter your name: ")

		if scanner.Scan() {
			name = strings.TrimSpace(scanner.Text())
		}

		if name != "" {
			break
		}

		fmt.Println("No name has been provided! Try again.")
	}

	fmt.Printf("Hello, %s! Welcome to Go!\n", name)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
		os.Exit(1)
	}
}