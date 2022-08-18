package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homepage).Methods("GET", "OPTIONS")                //Homepage
	router.HandleFunc("/list", getList).Methods("GET", "OPTIONS")             //View list
	router.HandleFunc("/list/{uid}", getItem).Methods("GET", "OPTIONS")       //View item
	router.HandleFunc("/list", createItem).Methods("POST", "OPTIONS")         //Add items
	router.HandleFunc("/list/{uid}", deleteItem).Methods("DELETE", "OPTIONS") //Delete item
	router.HandleFunc("/list", updateItem).Methods("PATCH", "OPTIONS")        //Update item

	log.Fatal(http.ListenAndServe(":8000", router))
}
