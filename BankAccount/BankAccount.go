package main

import (
	"fmt"
)

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

func (b *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		b.Balance += amount
		fmt.Printf("Deposited: $%.2f\n", amount)
	} else {
		fmt.Println("Invalid deposit amount.")
	}
}
func (b *BankAccount) Withdraw(amount float64) {
	if amount > 0 && amount <= b.Balance {
		b.Balance -= amount
		fmt.Printf("Withdrew: $%.2f\n", amount)
	} else {
		fmt.Println("Invalid or insufficient funds for withdrawal.")
	}
}

func (b *BankAccount) GetBalance() {
	fmt.Printf("Current Balance: $%.2f\n", b.Balance)
}

func Transaction(account *BankAccount, transactions []float64) {
	for _, t := range transactions {
		if t > 0 {
			account.Deposit(t)
		} else {
			account.Withdraw(-t)
		}
	}
}

func main() {
	account := BankAccount{
		AccountNumber: "12345",
		HolderName:    "John Doe",
		Balance:       0.0,
	}

	var choice int
	var amount float64

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Get Balance")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter deposit amount: ")
			fmt.Scan(&amount)
			account.Deposit(amount)
		case 2:
			fmt.Print("Enter withdrawal amount: ")
			fmt.Scan(&amount)
			account.Withdraw(amount)
		case 3:
			account.GetBalance()
		case 4:
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}
