package internal

import "time"

type Expense struct {
	ID 				int 		`json:"id"`
	Date 			time.Time 	`json:"date"`
	Amount  		float64 	`json:"amount"`
	Description 	string 		`json:"description"`
	Category 		string 		`json:"category"`
}