package main

import (
	"context"
	"log"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

collection := helper.ConnectDB()

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homepage).Methods("GET")                //Homepage
	router.HandleFunc("/list", getList).Methods("GET")             //View list
	router.HandleFunc("/list/{uid}", getItem).Methods("GET")       //View item
	router.HandleFunc("/list", createItem).Methods("POST")         //Add items
	router.HandleFunc("/list/{uid}", deleteItem).Methods("DELETE") //Delete item
	router.HandleFunc("/list", updateItem).Methods("PATCH")        //Update item

	log.Fatal(http.ListenAndServe(":8000", router))
}
