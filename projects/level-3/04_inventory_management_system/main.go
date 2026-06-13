package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var products []Product
	nextID := 1

	for {
		fmt.Println("\n=== Product Management System ===")
		fmt.Println("1. Add Product")
		fmt.Println("2. List Products")
		fmt.Println("3. Update Product")
		fmt.Println("4. Delete Product")
		fmt.Println("5. Search Product")
		fmt.Println("6. Adjust Stock")
		fmt.Println("0. Exit")

		option := readInt(scanner, "Choose an option: ")

		switch option {
		case 1:
			name := readString(scanner, "Name: ")

			price := readFloat(scanner, "Price: ")
			if price <= 0 {
				fmt.Println("Price must be greater than zero.")
				continue
			}

			quantity := readInt(scanner, "Quantity: ")
			if quantity < 0 {
				fmt.Println("Quantity cannot be negative.")
				continue
			}

			addProduct(&products, &nextID, name, price, quantity)

			fmt.Println("Product added successfully.")

		case 2:
			listProducts(products)

		case 3:
			id := readInt(scanner, "Enter product ID: ")

			name := readString(scanner, "New product name: ")

			price := readFloat(scanner, "New product price: ")
			if price <= 0 {
				fmt.Println("Price must be greater than zero.")
				continue
			}

			quantity := readInt(scanner, "New product quantity: ")
			if quantity < 0 {
				fmt.Println("Quantity cannot be negative.")
				continue
			}

			if updateProduct(
				products,
				id,
				name,
				price,
				quantity,
			) {
				fmt.Println("Product updated.")
			} else {
				fmt.Println("Product not found.")
			}

		case 4:
			id := readInt(scanner, "Enter product ID: ")

			if deleteProduct(&products, id) {
				fmt.Println("Product deleted.")
			} else {
				fmt.Println("Product not found.")
			}

		case 5:
			name := readString(scanner, "Enter product name: ")

			product, found := findProductByName(products, name)

			if found {
				fmt.Printf(
					"ID: %d | Name: %s | Price: %f | Quantity: %d\n",
					product.ID,
					product.Name,
					product.Price,
					product.Quantity,
				)
			} else {
				fmt.Println("Product not found.")
			}

		case 6:
			id := readInt(scanner, "\nEnter product ID: ")

			product, found := findProductById(products, id)

			if !found {
				fmt.Println("Product not found.")
				continue
			}

			fmt.Printf(
				"\nCurrent stock of %s: %d\n",
				product.Name,
				product.Quantity,
			)

			fmt.Println("\n1. Add Stock")
			fmt.Println("2. Remove Stock")

			option := readInt(scanner, "Choose option: ")

			amount := readInt(scanner, "Enter amount: ")

			if amount <= 0 {
				fmt.Println("Amount must be greater than zero.")
				continue
			}

			switch option {
			case 1:
				product.AddStock(amount)
				fmt.Println("Stock updated successfully.")

			case 2:
				err := product.RemoveStock(amount)

				if err != nil {
					fmt.Println(err)
					continue
				}

				fmt.Println("Stock updated successfully.")

			default:
				fmt.Println("Invalid option.")

			}

		case 0:
			fmt.Println("\nGoodbye! Closing system...")
			return

		default:
			fmt.Println("Invalid option.")

		}
	}

}

func addProduct(
	products *[]Product,
	nextID *int,
	name string,
	price float64,
	quantity int,
) {
	product := Product{
		ID:       *nextID,
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}

	*products = append(*products, product)

	*nextID++
}

func listProducts(products []Product) {
	if len(products) == 0 {
		fmt.Println("\nNo products found!")
		return
	}

	for _, product := range products {
		fmt.Printf(
			"ID: %d | Name: %s | Price: %.2f | Quantity: %d\n",
			product.ID,
			product.Name,
			product.Price,
			product.Quantity,
		)
	}
}

func updateProduct(products []Product, id int, name string, price float64, quantity int) bool {
	for i := range products {
		if products[i].ID == id {
			products[i].Name = name
			products[i].Price = price
			products[i].Quantity = quantity
			return true
		}
	}

	return false
}

func deleteProduct(products *[]Product, id int) bool {
	for i, product := range *products {
		if product.ID == id {
			*products = append(
				(*products)[:i],
				(*products)[i+1:]...,
			)

			return true
		}
	}

	return false
}

func findProductByName(products []Product, name string) (Product, bool) {
	for _, product := range products {
		if strings.EqualFold(product.Name, name) {
			return product, true
		}
	}
	return Product{}, false
}

func findProductById(products []Product, id int) (*Product, bool) {

	for i := range products {
		if products[i].ID == id {
			return &products[i], true
		}
	}

	return nil, false
}

func (p *Product) AddStock(amount int) {
	p.Quantity += amount
}

func (p *Product) RemoveStock(amount int) error {
	if amount > p.Quantity {
		return fmt.Errorf("Insufficient stock!")
	}

	p.Quantity -= amount
	return nil
}

func readString(scanner *bufio.Scanner, message string) string {
	for {
		fmt.Print(message)

		scanner.Scan()

		value := strings.TrimSpace(scanner.Text())

		if value != "" {
			return value
		}

		fmt.Println("Value cannot be empty! Try again.")
	}
}

func readFloat(scanner *bufio.Scanner, message string) float64 {
	for {
		fmt.Print(message)

		scanner.Scan()

		input := strings.TrimSpace(scanner.Text())

		value, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return value
		}

		fmt.Println("\nInvalid number! Try again.")
	}
}

func readInt(scanner *bufio.Scanner, message string) int {
	for {
		fmt.Print(message)

		scanner.Scan()

		input := strings.TrimSpace(scanner.Text())

		value, err := strconv.Atoi(input)
		if err == nil {
			return value
		}

		fmt.Println("\nInvalid number. Try again.")
	}
}
