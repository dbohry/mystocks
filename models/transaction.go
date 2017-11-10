package models

import "time"

// Transaction entity for transaction type
//
type Transaction struct {
	ID string `json:"id,omitempty"`
	User User `json:"user,omitempty"`
	Stock Stock `json:"stock,omitempty"`
	Type string `json:"type,omitempty"`
	Shares int `json:"shares,omitempty"`
	Price float64 `json:"price,omitempty"`
	Date time.Time `json:"date,omitempty"`
}