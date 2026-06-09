package main

import (
	"fmt"
	"strings"
)

func main() {
	var color string

	fmt.Println("\n=== Traffic Light Simulator ===")

	fmt.Print("Enter traffic light color (red, yellow or green): ")
	fmt.Scanln(&color)

	action := trafficLightAction(strings.ToLower(color))

	fmt.Printf("Action: %s\n", action)

}

func trafficLightAction(color string) string {
	switch color {
	case "red":
		return "Stop"
	case "yellow":
		return "Slow Down"
	case "green":
		return "Go"
	default:
		return "Invalid traffic light color"
	}
}
