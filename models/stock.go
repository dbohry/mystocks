package models

// Stock entity for stock type
//
type Stock struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Shares int32 `json:"shares,omitempty"`
}