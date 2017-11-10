package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/dbohry/mystocks/models"
	"github.com/dbohry/mystocks/services"
)

// GetTransactions return all transactions
//
func GetTransactions(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	result := services.GetTransactions(params["idUser"])
	json.NewEncoder(w).Encode(result)
}

// GetTransaction return a transaction by ID
//
func GetTransaction(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	result := services.GetTransactionByID(params["idUser"], params["id"])
	json.NewEncoder(w).Encode(result)
}

// SaveTransaction save a new transaction
//
func SaveTransaction(w http.ResponseWriter, req *http.Request) {
	var transaction models.Transaction
	_ = json.NewDecoder(req.Body).Decode(&transaction)
	saved := services.SaveTransaction(transaction)
	result := services.GetStockByID(saved.ID)
	json.NewEncoder(w).Encode(result)
}

// DeleteTransaction delete a transaction by ID
//
func DeleteTransaction(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	if services.DeleteTransaction(params["idUser"], params["id"]) {
		json.NewEncoder(w).Encode(params["id"] + " was removed!")
	} else {
		json.NewEncoder(w).Encode("Failed to remove " + params["id"])
	}
	
}