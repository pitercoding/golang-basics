package main

import "fmt"

func main() {
	var totalMinutes int

	fmt.Print("Enter total minutes: ")
	fmt.Scanln(&totalMinutes)

	hours, minutes := convertMinutes(totalMinutes)

	fmt.Printf(
		"%d minutes = %d hours and %d minutes\n",
		totalMinutes,
		hours,
		minutes,
	)
}

func convertMinutes(total int) (int, int) {
	hours := total / 60
	minutes := total % 60

	return hours, minutes
}
