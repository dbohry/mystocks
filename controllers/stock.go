package controllers

import (
	"encoding/json"
	"net/http"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"

	"github.com/dbohry/mystocks/models"
)

// GetStocks return all stocks
//
func GetStocks(w http.ResponseWriter, req *http.Request) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("mystocks").C("stocks")
	

	result := make([]models.Stock, 0, 10)
	err = c.Find(bson.M{}).All(&result)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(result)
}

// GetStock return a stocks by ID
//
func GetStock(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("mystocks").C("stocks")

	result := models.Stock{}
	err = c.Find(bson.M{"id": params["id"]}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(result)
}

// SaveStock save a new stock
//
func SaveStock(w http.ResponseWriter, req *http.Request) {	
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("mystocks").C("stocks")

	var stock models.Stock
	_ = json.NewDecoder(req.Body).Decode(&stock)
	err = c.Insert(&stock)
	if err != nil {
		log.Fatal(err)
	}

	result := models.Stock{}
	err = c.Find(bson.M{"id": &stock.ID}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(result)
}

// DeleteStock delete a stock by ID
//
func DeleteStock(w http.ResponseWriter, req *http.Request) {
}