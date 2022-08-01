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

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	shopping_list = append(shopping_list, item)

	json.NewEncoder(w).Encode(item)

	fmt.Println("createItem func called")
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	_deleteItemAtUid(params["uid"])

	json.NewEncoder(w).Encode(shopping_list)
}

func _deleteItemAtUid(uid string) {
	for index, item := range shopping_list {
		if item.UID == uid {
			shopping_list = append(shopping_list[:index], shopping_list[index+1:]...)
			break
		}
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homepage).Methods("GET")        //homepage
	router.HandleFunc("/list", getList).Methods("GET")     //view items in shopping list
	router.HandleFunc("/list", createItem).Methods("POST") //add items in shopping list
	router.HandleFunc("/list/{uid}", deleteItem).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	//append item to slice
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
