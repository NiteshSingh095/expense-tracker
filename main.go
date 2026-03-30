package main

import (
	"flag"
	"fmt"
	"os"
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

	fmt.Println("Amount:", *amount)
	fmt.Println("Description:", *description)
}
