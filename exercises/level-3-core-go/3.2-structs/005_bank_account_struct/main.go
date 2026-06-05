package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b BankAccount) Display() {
	fmt.Printf("Owner: %s\n", b.Owner)
	fmt.Printf("Balance: $%.2f\n", b.Balance)
}

func main() {
	var owner string
	var balance float64

	fmt.Println("\n=== Bank Account ===")

	fmt.Print("Enter account owner: ")
	fmt.Scanln(&owner)

	fmt.Print("Enter initial balance: ")
	fmt.Scanln(&balance)

	if balance < 0 {
		fmt.Println("Balance cannot be negative")
		return
	}

	account := BankAccount{
		Owner:   owner,
		Balance: balance,
	}

	fmt.Println("\n=== Account Information ===")
	account.Display()
}
