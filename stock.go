package main

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Stock entity for stock type
//
type Stock struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Shares string `json:"shares,omitempty"`
}

var stocks []Stock

// GetStocks return all stocks
//
func GetStocks(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(stocks)
}

// GetStock return a stocks by ID
//
func GetStock(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range stocks {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Stock{})
}

// SaveStock save a new stock
//
func SaveStock(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var stock Stock
	_ = json.NewDecoder(req.Body).Decode(&stock)
	stock.ID = params["id"]
	stocks = append(stocks, stock)
	json.NewEncoder(w).Encode(stocks)
}

// DeleteStock delete a stock by ID
//
func DeleteStock(w http.ResponseWriter, req *http.Request) {
	if req.URL == nil {
		fmt.Fprintf(w, "Please inform an ID")
	} else {
		fmt.Fprintf(w, "Stock removed: " + req.URL.Path[1:])
	}
}