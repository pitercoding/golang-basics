package main

import "fmt"

func main() {
	from := "min"
	to := "h"
	value := 120.0

	result, err := convert(value, from, to)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\nResult: %.2f %s = %.2f %s\n", value, from, result, to)
}

func convert(value float64, from, to string) (float64, error) {
	switch from + "-" + to {
	case "km-m":
		return value * 1000, nil
	case "m-km":
		return value / 1000, nil
	case "min-h":
		return value / 60, nil
	case "h-min":
		return value * 60, nil
	case "c-f":
		return (value * 1.8) + 32, nil
	case "f-c":
		return (value - 32) / 1.8, nil
	case "kg-g":
		return value * 1000, nil
	case "g-kg":
		return value / 1000, nil
	default:
		return 0, fmt.Errorf("Invalid conversion!")
	}
}
