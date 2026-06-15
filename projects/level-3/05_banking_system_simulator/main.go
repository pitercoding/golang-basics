package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BankAccount struct {
	ID      int
	Owner   string
	Balance float64
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var accounts []BankAccount
	nextID := 1

	for {
		fmt.Println("\n=== Banking System ===")
		fmt.Println("1. Create Account")
		fmt.Println("2. View Account")
		fmt.Println("3. Deposit")
		fmt.Println("4. Withdraw")
		fmt.Println("5. Transfer")
		fmt.Println("6. List Accounts")
		fmt.Println("0. Exit")

		option := readInt(scanner, "Choose an option: ")

		switch option {
		case 1:
			owner := readString(scanner, "Owner: ")

			createAccount(
				&accounts,
				&nextID,
				owner,
			)

			fmt.Println("Account created successfully.")

		case 2:
			id := readInt(scanner, "Enter account ID: ")

			account, found := findAccountByID(accounts, id)

			if !found {
				fmt.Printf("Account with id %d not found.\n", id)
				continue
			}

			fmt.Printf(
				"ID: %d | Owner: %s | Balance: %.2f\n",
				account.ID,
				account.Owner,
				account.Balance,
			)

		case 3:
			id := readInt(scanner, "Enter account ID: ")

			account, found := findAccountByID(accounts, id)

			if !found {
				fmt.Printf("Account with id %d not found.\n", id)
				continue
			}

			amount := readFloat(scanner, "Amount: ")

			if err := account.Deposit(amount); err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Printf("$%.2f deposited successfully!\n", amount)

		case 4:
			id := readInt(scanner, "Enter account ID: ")

			account, found := findAccountByID(accounts, id)

			if !found {
				fmt.Printf("Account with id %d not found.\n", id)
				continue
			}

			amount := readFloat(scanner, "Amount: ")

			if err := account.Withdraw(amount); err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Printf("$%.2f withdrawn successfully!\n", amount)

		case 5:
			fromID := readInt(scanner, "Enter source account ID (From): ")

			fromAccount, foundFrom := findAccountByID(accounts, fromID)

			if !foundFrom {
				fmt.Printf("Account with id %d not found.\n", fromID)
				continue
			}

			toID := readInt(scanner, "Enter destination account ID (To): ")

			toAccount, foundTo := findAccountByID(accounts, toID)

			if !foundTo {
				fmt.Printf("Account with id %d not found.\n", toID)
				continue
			}

			if fromID == toID {
				fmt.Println("Operation canceled: Source and destination accounts cannot be the same.")
				continue
			}

			amount := readFloat(scanner, "Amount to transfer: ")

			err := Transfer(fromAccount, toAccount, amount)
			if err != nil {
				fmt.Printf("Transfer failed: %v\n", err)
				continue
			}

			fmt.Printf("$%.2f transferred successfully from account %d to %d!\n", amount, fromID, toID)

		case 6:
			listAccounts(accounts)

		case 0:
			fmt.Println("\nGoodbye! Closing system...")
			return

		default:
			fmt.Println("Invalid option.")

		}

	}
}

func createAccount(
	accounts *[]BankAccount,
	nextID *int,
	owner string,
) {
	account := BankAccount{
		ID:      *nextID,
		Owner:   owner,
		Balance: 0,
	}

	*accounts = append(*accounts, account)

	*nextID++
}

func findAccountByID(accounts []BankAccount, id int) (*BankAccount, bool) {
	for i := range accounts {
		if accounts[i].ID == id {
			return &accounts[i], true
		}
	}

	return nil, false
}

func listAccounts(accounts []BankAccount) {
	if len(accounts) == 0 {
		fmt.Println("\nNo accounts found!")
		return
	}

	for _, account := range accounts {
		fmt.Printf(
			"ID: %d | Owner: %s | Balance: %.2f\n",
			account.ID,
			account.Owner,
			account.Balance,
		)
	}
}

func (b *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be greater than zero")
	}

	b.Balance += amount
	return nil
}

func (b *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be greater than zero")
	}

	if amount > b.Balance {
		return fmt.Errorf("Insufficient funds!")
	}

	b.Balance -= amount
	return nil
}

func Transfer(
	from *BankAccount,
	to *BankAccount,
	amount float64,
) error {

	if amount <= 0 {
		return fmt.Errorf("amount must be greater than zero")
	}

	if err := from.Withdraw(amount); err != nil {
		return err
	}

	if err := to.Deposit(amount); err != nil {
		return err
	}

	return nil
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

		fmt.Println("\nInvalid number. Try again.")
	}
}
