package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command: add, list, delete, summary")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		AddCommand()

	case "list":
		ListCommand()

	case "delete":
		DeleteCommand()

	case "summary":
		SummaryCommand()

	case "update":
		UpdateCommand()

	case "export":
		ExportToCSV()

	default:
		fmt.Println("Unknown command. Available commands: add, list, delete, summary")
	}
}