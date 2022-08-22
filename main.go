package main

import (
	"log"
	"net/http"

	"go-grocery-list-backend/middleware"
	"go-grocery-list-backend/routes"

	"github.com/gorilla/mux"
)

func main() {
	middleware.ConnectDB()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", routes.Homepage).Methods("GET")                //Homepage
	router.HandleFunc("/list", routes.GetList).Methods("GET")             //View list
	router.HandleFunc("/list/{uid}", routes.GetItem).Methods("GET")       //View item
	router.HandleFunc("/list", routes.CreateItem).Methods("POST")         //Add items
	router.HandleFunc("/list/{uid}", routes.DeleteItem).Methods("DELETE") //Delete item
	router.HandleFunc("/list", routes.UpdateItem).Methods("PUT")          //Update item

	log.Fatal(http.ListenAndServe(":8000", router))
}
