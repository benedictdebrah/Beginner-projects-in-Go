package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 0, fmt.Errorf("failed to read balance file: %w", err)
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse stored balance value: %w", err)
	}
	return balance, nil

}
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)

}

func deposit(amount float64) float64 {
	if amount > 0 {
		accountBalance += amount
		writeBalanceToFile(accountBalance)

		fmt.Printf("Successfully deposited: %.2f\n", amount)
	} else {
		fmt.Println("Amount must be greater than $0.00")
	}
	return accountBalance
}

func withdraw(amount float64) float64 {
	if amount > accountBalance {
		fmt.Printf("Insufficient funds. Your balance is: %.2f\n", accountBalance)
	} else if amount < 0 {
		fmt.Println("You cant withdraw negative amount ")

	} else {
		accountBalance -= amount
		writeBalanceToFile(accountBalance)
		fmt.Printf("Successfully withdrawn: %.2f\n", amount)
	}
	return accountBalance
}

func checkBalance() float64 {
	return accountBalance
}

var accountBalance, err = getBalanceFromFile()

func main() {
	if err != nil {
		fmt.Println("Error initializing account balance:", err)

	}

	fmt.Println("This is a bank app")

	for {
		fmt.Println("Please select your choice:")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Exit")

		fmt.Println("Please enter your choice:")
		var choice int
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			var depositAmount float64
			fmt.Printf("Enter amount to deposit:")
			fmt.Scanf("%f", &depositAmount)
			accountBalance = deposit(depositAmount)
		case 2:
			var withdrawAmount float64
			fmt.Println("Enter amount to withdraw:")
			fmt.Scanf("%f", &withdrawAmount)
			accountBalance = withdraw(withdrawAmount)
		case 3:
			balance := checkBalance()
			fmt.Printf("Your current balance is: %.2f\n", balance)
		case 4:
			fmt.Println("Exiting the application...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
