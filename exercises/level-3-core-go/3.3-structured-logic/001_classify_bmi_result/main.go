package main

import "fmt"

func main() {
	var weight, height float64

	fmt.Println("\n=== BMI Calculator ===")

	fmt.Print("Enter weight: ")
	fmt.Scanln(&weight)

	fmt.Print("Enter height: ")
	fmt.Scanln(&height)

	if weight <= 0 || height <= 0 {
		fmt.Println("Weight and height must be greater than zero!")
		return
	}

	bmi := calculateBMI(weight, height)
	classification := classifyBMI(bmi)

	fmt.Printf("\nBMI result: %.2f\n", bmi)
	fmt.Printf("BMI classification: %s\n", classification)
}

func calculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}

func classifyBMI(bmi float64) string {
	if bmi < 18.5 {
		return "Underweight"
	}

	if bmi < 25 {
		return "Normal Weight"
	}

	if bmi < 30 {
		return "Overweight"
	}

	return "Obesity"
}
