package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"go-grocery-list-backend/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

// var database := client.Database("Cluster0")
// var groceryCollection := database.Collection("grocery_list")
// var itemCollection := database.Collection("item")
// func getDatabase() (*mongo.Client, error) {
// 	if client == nil {
// 		err := godotenv.Load()

// 		if err != nil {
// 			log.Fatal("Error loading .env file")
// 			return nil, err
// 		}

// 		uri := os.Getenv("MONGODB_URL")
// 		c, err := mongo.NewClient(options.Client().ApplyURI(uri))

// 		if err != nil {
// 			log.Fatal(err)
// 			return nil, err
// 		}

// 		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 		err = client.Connect(ctx)

// 		if err != nil {
// 			return nil, err
// 		}

// 		client = c
// 	}

// 	return client
// }

func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint called: homepage()")

}

func GetList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var grocery_list []models.Item

	collection := client.Database("Cluster0").Collection("grocery_list")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var item models.Item
		cursor.Decode(&item)
		grocery_list = append(grocery_list, item)
	}

	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(w).Encode(grocery_list)

}

func GetItem(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	// itemName := mux.Vars(r)["uid"]

	// for _, singleItem := range grocery_list {
	// 	if singleItem.Name == itemName {
	// 		json.NewEncoder(w).Encode(singleItem)
	// 	}
	// }
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.Item

	json.NewDecoder(r.Body).Decode(&item)
	// c := getDatabase(client, error)
	collection := client.Database("Cluster0").Collection("item")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.InsertOne(ctx, item)

	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(result)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	// params := mux.Vars(r)

	// _deleteItemAtUid(params["uid"])

	// json.NewEncoder(w).Encode(grocery_list)
}

func _deleteItemAtUid(uid string) {
	// for index, item := range grocery_list {
	// 	if item.UID == uid {
	// 		grocery_list = append(grocery_list[:index], grocery_list[index+1:]...)
	// 		break
	// 	}
	// }
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	// var item Item
	// _ = json.NewDecoder(r.Body).Decode(&item)

	// //params := mux.Vars(r)["uid"]

	// for i, singleItem := range grocery_list {
	// 	if singleItem.UID == item.UID {
	// 		singleItem.Name = item.Name
	// 		singleItem.Price = item.Price
	// 		singleItem.Quantity = item.Quantity
	// 		grocery_list = append(grocery_list[:i], singleItem)
	// 		json.NewEncoder(w).Encode(singleItem)
	// 	}
	// }
}
