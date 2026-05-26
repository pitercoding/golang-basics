package main

import "fmt"

func main() {
	var month int

	fmt.Print("Enter month number (1-12): ")
	fmt.Scanln(&month)

	fmt.Println(getSeason(month))
}

func getSeason(month int) string {
	switch month {
	case 12, 1, 2:
		return "Winter"
	case 3, 4, 5:
		return "Spring"
	case 6, 7, 8:
		return "Summer"
	case 9, 10, 11:
		return "Autumn"
	default:
		return "Invalid month"
	}
}
