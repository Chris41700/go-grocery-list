package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func ConnectDB() *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb+srv://Chris41700:<TheLazyCoder41700>@cluster0.5k7ylkm.mongodb.net/?retryWrites=true&w=majority")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("grocerylist").Collection("items")

	return collection
}
