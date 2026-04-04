package main

import (
	"expense-tracker/internal"
	"flag"
	"fmt"
	"os"
	"time"
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
		listCommand()

	case "delete":
		deleteCommand()

	case "summary":
		summaryCommand()

	case "update":
		updateCommand()

	default:
		fmt.Println("Unknown command. Available commands: add, list, delete, summary")
	}
}

/// updateCommand handles the 'update' command to modify an existing expense by its ID.
func updateCommand() {
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)

	id := updateCmd.Int("id", 0, "Expense ID")
	desc := updateCmd.String("description", "", "New Description")
	amount := updateCmd.Float64("amount", 0, "New Amount")

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

/// summaryCommand handles the 'summary' command to calculate and display total expenses, optionally filtered by month.
func summaryCommand() {
	sumCmd := flag.NewFlagSet("summary", flag.ExitOnError)

	month := sumCmd.Int("month", 0, "Month(1-12)")

	sumCmd.Parse(os.Args[2:])

	if *month < 0 || *month > 12 {
		fmt.Println("Invalid Input. Month should be between 1 and 12.")
		return
	}

	expenses, err := internal.LoadExpense()
	if err != nil {
		fmt.Printf("Error loading expenses: %s\n", err)
		return
	}

	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	total := 0.0

	for _, expense := range expenses {
		if *month != 0 {
			if int(expense.Date.Month()) != *month {
				continue
			}
		}
		total += expense.Amount
	}

	if *month != 0 {
		fmt.Printf("Total expenses for month %d: $%0.2f\n", *month, total)
	} else {
		fmt.Printf("Total expenses: $%0.2f\n", total)
	}
}

// / deleteCommand handles the 'delete' command to remove an expense by its ID.
func deleteCommand() {
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

// / listCommand handles the 'list' command to display all expenses in a formatted manner.
func listCommand() {

	expenses, err := internal.LoadExpense()
	if err != nil {
		fmt.Printf("Error loading expenses: %s\n", err)
		return
	}

	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	fmt.Printf("%-5s %-12s %-15s %-10s\n", "ID", "Date", "Description", "Amount")

	for _, expense := range expenses {
		fmt.Printf("%-5d %-12s %-15s $%0.2f\n", expense.ID, expense.Date.Format("2006-01-02"), expense.Description, expense.Amount)
	}

}

// / addCommand handles the 'add' command to add a new expense with amount and description.
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
