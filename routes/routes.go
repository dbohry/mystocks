package routes

import (
	"fmt"
	"github.com/dbohry/mystocks/configs"
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
	router.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	router.HandleFunc("/users/{idUser}/transactions", controllers.GetTransactions).Methods("GET")
	router.HandleFunc("/users/{idUser}/transactions", controllers.SaveTransaction).Methods("POST")
	router.HandleFunc("/users/{idUser}/transactions/{id}", controllers.GetTransaction).Methods("GET")
	router.HandleFunc("/users/{idUser}/transactions/{id}", controllers.DeleteTransaction).Methods("DELETE")

	fmt.Println("Your server is running on " + configs.PORT)
	log.Fatal(http.ListenAndServe(configs.PORT, router))
}