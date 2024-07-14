package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	OrderID			primitive.ObjectID		`json:"id" bson:"_id,omitempty"`
	Room			primitive.ObjectID		`json:"roomId" bson:"roomId,omitempty"`
	Contact			string					`json:"contact"`
	CheckIn			string					`json:"checkIn"`
	CheckOut		string					`json:"checkOut"`
	OrderTime		string					`json:"orderTime"`
	TotalCharge		float64					`json:"totalCharge"`
}