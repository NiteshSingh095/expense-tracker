package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"expense-tracker/internal"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a command: add, list, delete, summary")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		addCommand()

	case "list":
		fmt.Println("List command triggered")

	case "delete":
		fmt.Println("Delete command triggered")

	case "summary":
		fmt.Println("Summary command triggered")

	default:
		fmt.Println("Unknown command. Available commands: add, list, delete, summary")
	}
}

/// addCommand handles the 'add' command to add a new expense with amount and description.
func addCommand() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	amount := addCmd.Float64("amount", 0.0, "Expense Amount")
	description := addCmd.String("description", "", "Expense Description")

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
	}

	expenses = append(expenses, newExpense)
	err = internal.SaveExpense(expenses)
	if err != nil {
		fmt.Printf("Error saving expenses: %s\n", err)
		return
	}

	fmt.Println("Expense added successfully!, ID : ", newExpense.ID)
}
