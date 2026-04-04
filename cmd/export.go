package cmd

import (
	"encoding/csv"
	"expense-tracker/internal"
	"fmt"
	"os"
	"strconv"

)

// / exportToCSV handles the 'export' command to export all expenses to a CSV file.
func ExportToCSV() {
	expenses, err := internal.LoadExpense()
	if err != nil {
		fmt.Printf("Error loading expenses: %s\n", err)
		return
	}

	file, err := os.Create("expenses.csv")
	if err != nil {
		fmt.Printf("Error creating CSV file: %s\n", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Date", "Description", "Amount", "Category"})

	for _, expense := range expenses {
		writer.Write([]string{
			strconv.Itoa(expense.ID),
			expense.Date.Format("2006-01-02"),
			expense.Description,
			fmt.Sprintf("%.2f", expense.Amount),
			expense.Category,
		})
	}
}