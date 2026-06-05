package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

func (p Product) Display() {
	fmt.Printf("%s - $%.2f\n", p.Name, p.Price)
}

func main() {
	var name string
	var price float64
	
	fmt.Println("\n=== Product Struct ===")
	
	fmt.Print("Enter product name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter product price: ")
	fmt.Scanln(&price)

	if price < 0 {
		fmt.Println("Price cannot be negative")
		return
	}

	product := Product {
		Name: name,
		Price: price,
	}

	fmt.Println("\n--- Product ---")
	product.Display()
}