package models

// Stock entity for stock type
//
type Stock struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Shares string `json:"shares,omitempty"`
}

var Stocks []Stock

//mock
//
func init() {
	Stocks = append(Stocks, Stock{ID: "WEGE3", Name: "WEG"})
	Stocks = append(Stocks, Stock{ID: "ITUB3", Name: "Itaú Unibanco"})
	Stocks = append(Stocks, Stock{ID: "LREN3", Name: "Lojas Renner"})
}