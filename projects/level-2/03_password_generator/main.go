package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\n=== Password Generator ===")

	passLength := readInt(scanner, "Enter password length: ")

	if passLength <= 0 {
		fmt.Println("Invalid length")
		return
	}

	password := generatePassword(passLength)

	fmt.Println("\nGenerated password:", password)

}

func generatePassword(passLength int) string {
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	symbols := "!@#$%^&*"

	all := lower + upper + numbers + symbols

	rand.Seed(time.Now().UnixNano())

	password := strings.Builder{}

	for i := 0; i < passLength; i++ {
		index := rand.Intn(len(all))
		password.WriteByte(all[index])
	}

	return password.String()
}

func readInt(scanner *bufio.Scanner, message string) int {
	for {
		fmt.Print(message)

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		value, err := strconv.Atoi(input)
		if err == nil {
			return value
		}

		fmt.Println("Invalid input. Try again.")
	}
}
