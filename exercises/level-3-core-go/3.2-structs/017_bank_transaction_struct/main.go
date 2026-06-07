package main

import "fmt"

type Transaction struct {
	Type   string
	Amount float64
}

type BankAccount struct {
	Owner        string
	Balance      float64
	Transactions []Transaction
}

func (b *BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Invalid deposit amount.")
		return
	}

	b.Balance += amount

	b.Transactions = append(
		b.Transactions,
		Transaction{
			Type:   "Deposit",
			Amount: amount,
		},
	)
}

func (b *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("invalid withdrawal amount")
	}

	if amount > b.Balance {
		return fmt.Errorf("insufficient funds")
	}

	b.Balance -= amount

	b.Transactions = append(
		b.Transactions,
		Transaction{
			Type:   "Withdraw",
			Amount: amount,
		},
	)

	return nil
}

func (b BankAccount) ShowTransactions() {
	if len(b.Transactions) == 0 {
		fmt.Println("No transactions found.")
		return
	}

	fmt.Println("\n=== Transaction History ===")

	for i, transaction := range b.Transactions {
		fmt.Printf(
			"%d. %s - $%.2f\n",
			i+1,
			transaction.Type,
			transaction.Amount,
		)
	}
}

func (b BankAccount) TotalTransactions() int {
	return len(b.Transactions)
}

func main() {
	account := BankAccount{
		Owner: "Bia",
	}

	account.Deposit(1000)
	account.Deposit(500)

	if err := account.Withdraw(300); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Owner: %s\n", account.Owner)
	fmt.Printf("Balance: $%.2f\n", account.Balance)

	account.ShowTransactions()

	fmt.Printf(
	"Total transactions: %d\n",
	account.TotalTransactions(),
)
}
