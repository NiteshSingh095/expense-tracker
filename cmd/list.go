package cmd

import (
	"expense-tracker/internal"
	"fmt"
)

/// listCommand handles the 'list' command to display all expenses in a formatted manner.
func ListCommand() {

	expenses, err := internal.LoadExpense()
	if err != nil {
		fmt.Printf("Error loading expenses: %s\n", err)
		return
	}

	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	fmt.Printf("%-5s %-12s %-15s %-10s %-10s\n", "ID", "Date", "Desc", "Amount", "Category")

	for _, expense := range expenses {
		fmt.Printf("%-5d %-12s %-15s $%-10.2f %-10s\n", expense.ID, expense.Date.Format("2006-01-02"), expense.Description, expense.Amount, expense.Category)
	}
}