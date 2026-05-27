package main

import (
	"fmt"
)

func main() {
	var weight float64
	var height float64

	fmt.Print("Enter weight (kg): ")
	fmt.Scanln(&weight)

	fmt.Print("Enter height (m): ")
	fmt.Scanln(&height)

	bmi := calculateBMI(weight, height)

	fmt.Printf("BMI: %.2f\n", bmi)

	if bmi < 18.5 {
		fmt.Println("Underweight")
	} else if bmi < 25 {
		fmt.Println("Normal weight")
	}

}

func calculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}
