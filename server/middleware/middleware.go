package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go.mongodb.org/driver-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
)

const connectionString = "Connection String"

const dbName = "Cluster0"

const collName = "grocerylist"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection instance created!")
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint called: homepage()")
}

func getList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(grocery_list)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	itemUID := mux.Vars(r)["uid"]

	for _, singleItem := range grocery_list {
		if singleItem.UID == itemUID {
			json.NewEncoder(w).Encode(singleItem)
		}
	}
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
