package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type GroceryList struct {
	ID       primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty"`
	Item     *Item              `json:"Name, omitempty" bson:"Name, omitempty"`
	Total    float64            `json:"Total, omitempty" bson:"Total, omitempty"`
	Quantity int                `json:"Quantity, omitempty" bson:"Quantity, omitempty"`
}

type Item struct {
	Name     string  `json:"Name, omitempty" bson:"Name, omitempty"`
	Price    float64 `json:"Price, omitempty" bson:"Price, omitempty"`
	Quantity int     `json:"Quantity, omitempty", bson:"Quantity, omitempty"`
}
