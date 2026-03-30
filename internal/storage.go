package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

/// LoadExpense reads the expenses from the JSON file and returns a slice of Expense structs.
func LoadExpense() ([]Expense, error) {
	file, err := os.ReadFile("data/expense.json")
	if err != nil {
		return nil, fmt.Errorf("Error occured while reading file : %s", err)
	}

	var expenses []Expense
	err = json.Unmarshal(file, &expenses)
	if err != nil {
		return nil, fmt.Errorf("Error occured while unmarshalling JSON : %s", err)
	}
	return expenses, nil
}

/// SaveExpense takes a slice of Expense structs and saves it to the JSON file.
func SaveExpense(expense []Expense) error {
	data, err := json.MarshalIndent(expense, "", "  ")
	if err != nil {
		return fmt.Errorf("Error occured while marshalling JSON : %s", err)
	}

	return os.WriteFile("data/expense.json", data, 0644)
}

/// GetNextId generates the next unique ID for a new expense based on the existing expenses.
func GetNextId(expense []Expense) int {
	if len(expense) == 0 {
		return 1
	}

	return expense[len(expense)-1].ID + 1
}