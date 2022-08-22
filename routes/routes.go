package routes

import (
	"encoding/json"
	"fmt"
	"go-grocery-list-backend/models"
	"net/http"

	"github.com/gorilla/mux"
)

var grocery_list []models.Item

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint called: homepage()")

}

func getList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var grocery_list models.GroceryList

	json.NewEncoder(w).Encode(grocery_list)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	itemName := mux.Vars(r)["name"]

	for _, singleItem := range grocery_list {
		if singleItem.Name == itemName {
			json.NewEncoder(w).Encode(singleItem)
		}
	}
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	models.GroceryList = append(grocery_list, item)

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

	//params := mux.Vars(r)["uid"]

	for i, singleItem := range grocery_list {
		if singleItem.UID == item.UID {
			singleItem.Name = item.Name
			singleItem.Price = item.Price
			singleItem.Quantity = item.Quantity
			grocery_list = append(grocery_list[:i], singleItem)
			json.NewEncoder(w).Encode(singleItem)
		}
	}
}
