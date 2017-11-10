package models

import "time"

// Transaction entity for transaction type
//
type Transaction struct {
	ID string `json:"id,omitempty"`
	User *User `json:"user,omitempty"`
	Type string `json:"type,omitempty"`
	Shares int32 `json:"shares,omitempty"`
	Price float32 `json:"price,omitempty"`
	Date time.Time `json:"date,omitempty"`
}