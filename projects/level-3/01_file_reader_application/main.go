package main

import (
	"fmt"
	"go/constant"
	"os"
)

func main() {
	var filePath string

	fmt.Println("\n=== File Reader ===")

	fmt.Print("Enter file path: ")
	fmt.Scanln(&filePath)

	content, err := readFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("\n--- File Content ---")
	fmt.Println(content)
}

func readFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(content), nil
}