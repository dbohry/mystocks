package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/dbohry/mystocks/models"
	"github.com/dbohry/mystocks/services"
)

// GetUsers return all users
//
func GetUsers(w http.ResponseWriter, req *http.Request) {
	result := services.GetUsers()
	json.NewEncoder(w).Encode(result)
}

// GetUserByName return an user by name
//
func GetUser(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	result := services.GetUserById(params["id"])
	json.NewEncoder(w).Encode(result)
}

// SaveUser save a new user
//
func SaveUser(w http.ResponseWriter, req *http.Request) {
	var user models.User
	_ = json.NewDecoder(req.Body).Decode(&user)
	saved := services.SaveUser(user)
	result := services.GetUserById(saved.User)
	json.NewEncoder(w).Encode(result)
}

// DeleteUser delete an user by ID
//
func DeleteUser(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	if services.DeleteStock(params["id"]) {
		json.NewEncoder(w).Encode(params["id"] + " was removed!")
	} else {
		json.NewEncoder(w).Encode("Failed to remove " + params["id"])
	}
	
}