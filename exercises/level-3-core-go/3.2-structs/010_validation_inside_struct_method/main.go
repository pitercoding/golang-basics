package main

import (
	"fmt"
)

type Car struct {
	Brand    string
	Model    string
	Position int
}

func (c *Car) Move(distance int) error {
	if distance < 0 {
		return fmt.Errorf("Distance cannot be negative!")
	}

	c.Position += distance
	return nil
}

func (c *Car) Reverse(distance int) error {
	c.Position -= distance

	if c.Position < 0 {
		c.Position = 0
	}

	if distance < 0 {
		return fmt.Errorf("Distance cannot be negative!")
	}
	return nil
}

func (c Car) ShowPosition() {
	fmt.Printf("Current position: %d\n", c.Position)
}

func main() {

	car := Car{
		Brand:    "Toyota",
		Model:    "Corolla",
		Position: 0,
	}

	fmt.Println("\n=== Validation Inside Struct Method ===")

	car.ShowPosition()

	if err := car.Move(10); err != nil {
		fmt.Println(err)
	}

	car.ShowPosition()

	if err := car.Move(-5); err != nil {
		fmt.Println(err)
	}

	car.ShowPosition()

	if err := car.Reverse(3); err != nil {
		fmt.Println(err)
	}

	car.ShowPosition()
}
