package main

import (
	"log"
	"net/http"

	"github.com/dbohry/mystocks/controllers"
	"github.com/gorilla/mux"
)

//Routes create app URL routes
//
func init() {
	router := mux.NewRouter()
	
		router.HandleFunc("/stocks", controllers.GetStocks).Methods("GET")
		router.HandleFunc("/stocks", controllers.SaveStock).Methods("POST")
		router.HandleFunc("/stocks/{id}", controllers.GetStock).Methods("GET")
		router.HandleFunc("/stocks/{id}", controllers.DeleteStock).Methods("DELETE")
		
	log.Fatal(http.ListenAndServe(":8080", router))
}