package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

type Inventory struct {
	Products []Product
}

func (i *Inventory) AddProduct(product Product) {
	i.Products = append(i.Products, product)
}

func (i Inventory) ListProducts() {
	if len(i.Products) == 0 {
		fmt.Println("No products found.")
		return
	}

	fmt.Println("\n=== Products ===")

	for index, product := range i.Products {
		fmt.Printf("%d. %s\n", index+1, product.Name)
		fmt.Printf("   Price: $%.2f\n", product.Price)
		fmt.Printf("   Quantity: %d\n\n", product.Quantity)
	}
}

func (i Inventory) TotalValue() float64 {
	var total float64

	for _, product := range i.Products {
		total += product.Price * float64(product.Quantity)
	}

	return total
}

func (i Inventory) TotalProducts() int {
	total := 0

	for _, product := range i.Products {
		total += product.Quantity
	}

	return total
}

func main() {
	inventory := Inventory{}

	inventory.AddProduct(Product{
		Name:     "Notebook",
		Price:    3500,
		Quantity: 2,
	})

	inventory.AddProduct(Product{
		Name:     "Mouse",
		Price:    120,
		Quantity: 5,
	})

	inventory.ListProducts()

	fmt.Printf(
		"Total inventory value: $%.2f\n",
		inventory.TotalValue(),
	)

	fmt.Printf(
		"Total products: %d\n",
		inventory.TotalProducts(),
	)
}
