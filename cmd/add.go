package cmd

import (
	"expense-tracker/internal"
	"flag"
	"fmt"
	"os"
	"time"
)

// / addCommand handles the 'add' command to add a new expense with amount and description.
func AddCommand() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	amount := addCmd.Float64("amount", 0.0, "Expense Amount")
	description := addCmd.String("description", "", "Expense Description")
	category := addCmd.String("category", "general", "Expense Category")

	addCmd.Parse(os.Args[2:])

	if *description == "" || *amount <= 0 {
		fmt.Println("Invalid Input. Description and Positive Amount are required.")
		return
	}

	expenses, err := internal.LoadExpense()
	if err != nil {
		fmt.Printf("Error loading expenses: %s\n", err)
		return
	}

	newExpense := internal.Expense{
		ID:          internal.GetNextId(expenses),
		Date:        time.Now(),
		Amount:      *amount,
		Description: *description,
		Category:    *category,
	}

	expenses = append(expenses, newExpense)
	err = internal.SaveExpense(expenses)
	if err != nil {
		fmt.Printf("Error saving expenses: %s\n", err)
		return
	}

	fmt.Println("Expense added successfully!, ID : ", newExpense.ID)
}