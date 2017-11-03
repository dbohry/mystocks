package routes

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/dbohry/mystocks/controllers"
)

//Start create routes and start web server
//
func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/stocks", controllers.GetStocks).Methods("GET")
	router.HandleFunc("/stocks", controllers.SaveStock).Methods("POST")
	router.HandleFunc("/stocks/{id}", controllers.GetStock).Methods("GET")
	router.HandleFunc("/stocks/{id}", controllers.DeleteStock).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":8080", router))
}