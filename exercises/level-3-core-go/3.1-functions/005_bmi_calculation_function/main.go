package main

import "fmt"

func main() {
	var weight, height float64

	fmt.Println("\n=== BMI Calculator ===")

	fmt.Print("Enter weight (kg): ")
	fmt.Scanln(&weight)

	fmt.Print("Enter height (m): ")
	fmt.Scanln(&height)

	if weight <= 0 || height <= 0 {
		fmt.Println("Invalid values: weight and height must be greater than 0")
		return
	}

	bmi := calculateBMI(weight, height)
	category := classifyBMI(bmi)

	fmt.Printf("\nBMI: %.2f\n", bmi)
	fmt.Printf("Category: %s\n", category)
}

func calculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}

func classifyBMI(bmi float64) string {
	if bmi < 18.5 {
		return "Underweight"
	}
	if bmi < 25 {
		return "Normal"
	}
	if bmi < 30 {
		return "Overweight"
	}
	return "Obesity"
}
