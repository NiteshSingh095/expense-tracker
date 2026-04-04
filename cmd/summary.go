package cmd

import(
	"expense-tracker/internal"
	"flag"
	"fmt"
	"os"
)

/// summaryCommand handles the 'summary' command to calculate and display total expenses, optionally filtered by month.
func SummaryCommand() {
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