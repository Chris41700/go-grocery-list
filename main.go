package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	UID      string  `json:"UID"`
	Name     string  `json:"Name"`
	Price    float64 `json:"Price"`
	Quantity int     `json:"Quantity"`
}

var grocery_list []Item

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint called: homepage()")
}

func getList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Function Called: getList()")

	json.NewEncoder(w).Encode(grocery_list)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	grocery_list = append(grocery_list, item)

	json.NewEncoder(w).Encode(item)

	fmt.Println("createItem func called")
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	_deleteItemAtUid(params["uid"])

	json.NewEncoder(w).Encode(grocery_list)
}

func _deleteItemAtUid(uid string) {
	for index, item := range grocery_list {
		if item.UID == uid {
			grocery_list = append(grocery_list[:index], grocery_list[index+1:]...)
			break
		}
	}
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	params := mux.Vars(r)

	//Deletes the item at UID
	_deleteItemAtUid(params["uid"])

	//Create item with new data
	grocery_list = append(grocery_list, item)

	json.NewEncoder(w).Encode(grocery_list)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homepage).Methods("GET")                //homepage
	router.HandleFunc("/list", getList).Methods("GET")             //view items
	router.HandleFunc("/list", createItem).Methods("POST")         //add items
	router.HandleFunc("/list/{uid}", deleteItem).Methods("DELETE") //delete item
	router.HandleFunc("/list/{uid}", updateItem).Methods("PUT")    //update item

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	//append item to slice
	grocery_list = append(grocery_list, Item{
		UID:      "0",
		Name:     "Cheese",
		Price:    4.99,
		Quantity: 1,
	})

	grocery_list = append(grocery_list, Item{
		UID:      "1",
		Name:     "Milk",
		Price:    3.25,
		Quantity: 2,
	})

	handleRequests()
}
