package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Car struct {
	Brand    string
	Model    string
	Position int
}

func (c *Car) Move(distance int) {
	c.Position += distance
}

func (c *Car) Reverse(distance int) {
	c.Position -= distance

	if c.Position < 0 {
		c.Position = 0
	}
}

func (c Car) ShowPosition() {
	fmt.Printf("Current position: %d\n", c.Position)
}

func ReadInt(scanner *bufio.Scanner, message string) int {
	for {
		fmt.Print(message)

		if !scanner.Scan() {
			fmt.Println("Failed to read input.")
			os.Exit(1)
		}

		input := strings.TrimSpace(scanner.Text())

		value, err := strconv.Atoi(input)
		if err == nil {
			return value
		}

		fmt.Println("Invalid input. Try again.")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	car := Car{
		Brand:    "Toyota",
		Model:    "Corolla",
		Position: 0,
	}

	for {

		fmt.Println("\n=== Move Car ===")
		fmt.Println("1. Move")
		fmt.Println("2. Reverse")
		fmt.Println("0. Exit")

		option := ReadInt(scanner, "Choose an option: ")

		switch option {
		case 1:
			distance := ReadInt(scanner, "Enter distance to move: ")

			if distance < 0 {
				fmt.Println("Distance cannot be negative")
				continue
			}

			car.Move(distance)
			car.ShowPosition()

		case 2:
			distance := ReadInt(scanner, "Enter distance to reverse: ")

			if distance < 0 {
				fmt.Println("Distance cannot be negative")
				continue
			}

			car.Reverse(distance)
			car.ShowPosition()

		case 0:
			fmt.Println("\nThank you! Exiting...")
			return

		default:
			fmt.Println("Invalid option! Try again.")
		}

	}

}
