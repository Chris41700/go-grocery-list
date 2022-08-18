package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"go-grocery-list-backend/models"
	"log"
	"net/http"
	"os"

	"../models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB() *mongo.Collection {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URL")

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("grocerylist").Collection("items")

	return collection
}

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
