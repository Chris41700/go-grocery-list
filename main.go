package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	UID   string  `json:"UID"`
	Name  string  `json:"Name"`
	Price float64 `json:"Price"`
}

var shopping_list []Item

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint called: homepage()")
}

func getList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Function Called: getList()")

	json.NewEncoder(w).Encode(shopping_list)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homepage).Methods("GET")    //Homepage
	router.HandleFunc("/list", getList).Methods("GET") //View items in shopping list

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	//Append item to slice
	shopping_list = append(shopping_list, Item{
		UID:   "0",
		Name:  "Cheese",
		Price: 4.99,
	})

	shopping_list = append(shopping_list, Item{
		UID:   "1",
		Name:  "Milk",
		Price: 3.25,
	})

	handleRequests()
}
