package main

import (
	"fmt"
)

func main() {
	scores := map[string]float64{
		"John":  90,
		"Mary":  70,
		"Peter": 85,
		"Anna":  60,
	}

	fmt.Println("\n=== Students / Grades ===")
	for student, grade := range scores {
		fmt.Printf("%s: %.2f\n", student, grade)
	}

	sum, avg, max, min := computeStats(scores)

	fmt.Println("\n=== Stats ===")
	fmt.Printf("Sum: %.2f | Avg: %.2f | Max: %.2f | Min: %.2f\n", sum, avg, max, min)

}

func computeStats(scores map[string]float64) (sum, avg, max, min float64){
	sum = 0.0
	count := 0
	first := true

	for _, grade := range scores {
		sum += grade
		count++

		if first {
			max = grade
			min = grade
			first = false
		} 
		
		if grade > max {
			max = grade
		} 
		
		if grade < min {
			min = grade
		}
	}

	avg = sum / float64(count)

	return sum, avg, max, min
}