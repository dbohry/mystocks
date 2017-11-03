package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/dbohry/mystocks/models"
	"github.com/dbohry/mystocks/services"
)

const collection = "stocks"

// GetStocks return all stocks
//
func GetStocks(w http.ResponseWriter, req *http.Request) {
	result := services.GetStocks()
	json.NewEncoder(w).Encode(result)
}

// GetStock return a stocks by ID
//
func GetStock(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	result := services.GetStockByID(params["id"])
	json.NewEncoder(w).Encode(result)
}

// SaveStock save a new stock
//
func SaveStock(w http.ResponseWriter, req *http.Request) {
	var stock models.Stock
	_ = json.NewDecoder(req.Body).Decode(&stock)
	saved := services.CreateStock(stock)
	result := services.GetStockByID(saved.ID)
	json.NewEncoder(w).Encode(result)
}

// DeleteStock delete a stock by ID
//
func DeleteStock(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	services.DeleteStock(params["id"])
	json.NewEncoder(w).Encode(params["id"] + " was removed!")
}