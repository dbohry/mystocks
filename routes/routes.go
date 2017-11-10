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
	
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users", controllers.SaveUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.GetUserByName).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	router.HandleFunc("/users/{idUser}/transactions", controllers.GetTransactions).Methods("GET")
	router.HandleFunc("/users/{idUser}/transactions", controllers.SaveTransaction).Methods("POST")
	router.HandleFunc("/users/{idUser}/transactions/{id}", controllers.GetTransaction).Methods("GET")
	router.HandleFunc("/users/{idUser}/transactions/{id}", controllers.DeleteTransaction).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}