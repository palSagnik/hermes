package models

import "time"

// transactions table
// id, users, splittype, total

// Each expense would have a user
// The total amount in a transaction
type Expense struct {
	ExpenseID     int       `json:"expenseid"`
	UserID        int 	    `json:"userid"`
	User		  User      `json:"user"`	
	Amount        float64   `json:"total"`
	Description   string    `json:"desc"`
	CreatedAt 	  time.Time     			
}