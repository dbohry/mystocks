package controllers

import (
	"encoding/json"
	"net/http"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"

	"github.com/dbohry/mystocks/models"
	"github.com/dbohry/mystocks/configs"
	"github.com/dbohry/mystocks/tools"
)

const collection = "stocks"

// GetStocks return all stocks
//
func GetStocks(w http.ResponseWriter, req *http.Request) {
	session, err := mgo.Dial(configs.HOST)
	tools.ValidatePanic(err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(configs.DB).C(collection)

	result := make([]models.Stock, 0, 10)
	err = c.Find(bson.M{}).All(&result)
	tools.ValidateFatal(err)

	json.NewEncoder(w).Encode(result)
}

// GetStock return a stocks by ID
//
func GetStock(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	
	session, err := mgo.Dial(configs.HOST)
	tools.ValidatePanic(err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(configs.DB).C(collection)

	result := models.Stock{}
	err = c.Find(bson.M{"id": params["id"]}).One(&result)
	tools.ValidateFatal(err)

	json.NewEncoder(w).Encode(result)
}

// SaveStock save a new stock
//
func SaveStock(w http.ResponseWriter, req *http.Request) {	
	session, err := mgo.Dial(configs.HOST)
	tools.ValidatePanic(err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(configs.DB).C(collection)

	var stock models.Stock
	_ = json.NewDecoder(req.Body).Decode(&stock)
	err = c.Insert(&stock)
	tools.ValidateFatal(err)

	result := models.Stock{}
	err = c.Find(bson.M{"id": &stock.ID}).One(&result)
	tools.ValidateFatal(err)

	json.NewEncoder(w).Encode(result)
}

// DeleteStock delete a stock by ID
//
func DeleteStock(w http.ResponseWriter, req *http.Request) {
}