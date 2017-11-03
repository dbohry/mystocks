package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dbohry/mystocks/models"
	"github.com/gorilla/mux"
)

var stocks = models.Stocks

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
	json.NewEncoder(w).Encode(&models.Stock{})
}

// SaveStock save a new stock
//
func SaveStock(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var stock models.Stock
	_ = json.NewDecoder(req.Body).Decode(&stock)
	stock.ID = params["id"]
	stocks = append(stocks, stock)
	json.NewEncoder(w).Encode(stocks)
}

// DeleteStock delete a stock by ID
//
func DeleteStock(w http.ResponseWriter, req *http.Request) {
}