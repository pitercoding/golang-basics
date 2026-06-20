package main

import "fmt"

type Person struct {
	Name string
	City string
}

func main() {
	people := []Person{
		{"John", "Berlin"},
		{"Mary", "Paris"},
		{"Peter", "Berlin"},
		{"Anna", "Madrid"},
		{"Bob", "Paris"},
	}

	result := groupByCity(people)

	fmt.Println(result)
}

func groupByCity(people []Person) map[string][]Person {
	result := make(map[string][]Person)

	for _, person := range people {
		result[person.City] = append(
			result[person.City],
			person,
		)
	}

	return result
}
