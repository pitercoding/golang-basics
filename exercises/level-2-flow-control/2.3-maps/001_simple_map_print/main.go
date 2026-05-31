package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "John",
		"city":    "Berlin",
		"country": "Germany",
	}

	fmt.Println("Name:", person["name"])
	fmt.Println("City:", person["city"])
	fmt.Println("Country:", person["country"])
}
