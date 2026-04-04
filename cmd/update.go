package cmd

import (
	"expense-tracker/internal"
	"flag"
	"fmt"
	"os"
)

/// updateCommand handles the 'update' command to modify an existing expense by its ID.
func UpdateCommand() {
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)

	id := updateCmd.Int("id", 0, "Expense ID")
	desc := updateCmd.String("description", "", "New Description")
	amount := updateCmd.Float64("amount", 0, "New Amount")
	catergory := updateCmd.String("category", "general", "New Category")

	updateCmd.Parse(os.Args[2:])

	if *id <= 0 {
		fmt.Println("Invalid Input. Positive ID is required.")
		return
	}

	expenses, err := internal.LoadExpense()
	if err != nil {
		fmt.Printf("Error loading expenses: %s\n", err)
		return
	}

	found := false

	for index, expense := range expenses {
		if expense.ID == *id {
			found = true

			if *desc != "" {
				expenses[index].Description = *desc
			}

			if *amount > 0 {
				expenses[index].Amount = *amount
			}

			expenses[index].Category = *catergory

			break
		}
	}

	if !found {
		fmt.Printf("Expense with ID %d not found.\n", *id)
		return
	}

	err = internal.SaveExpense(expenses)
	if err != nil {
		fmt.Printf("Error saving expenses: %s\n", err)
		return
	}

	fmt.Printf("Expense with ID %d updated successfully!\n", *id)
}