package cmd

import (
	"expense-tracker/internal"
	"flag"
	"fmt"
	"os"
)

/// deleteCommand handles the 'delete' command to remove an expense by its ID.
func DeleteCommand() {
	delCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	id := delCmd.Int("id", 0, "Expense ID")

	delCmd.Parse(os.Args[2:])

	if *id <= 0 {
		fmt.Println("Invalid Input. Positive ID is required.")
		return
	}

	expenses, err := internal.LoadExpense()
	if err != nil {
		fmt.Printf("Error loading expenses: %s\n", err)
		return
	}

	var updatedExpenses []internal.Expense

	found := false
	for _, expense := range expenses {
		if expense.ID == *id {
			found = true
			continue
		}
		updatedExpenses = append(updatedExpenses, expense)
	}

	if !found {
		fmt.Printf("Expense with ID %d not found.\n", *id)
		return
	}

	err = internal.SaveExpense(updatedExpenses)
	if err != nil {
		fmt.Printf("Error saving expenses: %s\n", err)
		return
	}

	fmt.Printf("Expense with ID %d deleted successfully!\n", *id)
}