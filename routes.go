package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Routes create app URL routes
//
func init() {
	router := mux.NewRouter()
	
		stocks = append(stocks, Stock{ID: "WEGE3", Name: "WEG"})
		stocks = append(stocks, Stock{ID: "ITUB3", Name: "Itaú Unibanco"})
		stocks = append(stocks, Stock{ID: "LREN3", Name: "Lojas Renner"})
	
		router.HandleFunc("/stock", GetStocks).Methods("GET")
		router.HandleFunc("/stock/{id}", GetStock).Methods("GET")
		router.HandleFunc("/stock/{id}", SaveStock).Methods("POST")
		router.HandleFunc("/stock/{id}", DeleteStock).Methods("DELETE")
		
	log.Fatal(http.ListenAndServe(":8080", router))
}