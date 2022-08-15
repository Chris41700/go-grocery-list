package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type GroceryList struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `json:"Name"`
	Price    float64            `json:"Price"`
	Quantity int                `json:"Quantity"`
}
