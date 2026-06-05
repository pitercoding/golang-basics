package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

func (p Product) ApplyDiscount(percent float64) float64 {
	return p.Price - (p.Price * percent / 100)
}

func main() {
	var name string
	var price, discount float64

	fmt.Println("\n=== Product Discount Calculator ===")

	fmt.Print("Enter product Name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter Product Price: ")
	fmt.Scanln(&price)

	if price < 0 {
		fmt.Println("Price cannot be negative!")
		return
	}

	fmt.Print("Enter discount percentage: ")
	fmt.Scanln(&discount)

	if discount < 0 || discount > 100 {
		fmt.Println("Discount must be between 0 and 100")
		return
	}

	product := Product{
		Name:  name,
		Price: price,
	}

	finalPrice := product.ApplyDiscount(discount)

	fmt.Printf("\nProduct: %s\n", product.Name)
	fmt.Printf("Original Price: $%.2f\n", product.Price)
	fmt.Printf("Discount: %.2f%%\n", discount)
	fmt.Printf("Final Price: $%.2f\n", finalPrice)
}
