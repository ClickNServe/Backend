package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bed struct {
	BedID		primitive.ObjectID 		`json:"id" bson:"_id,omitempty"`
	BedType		string					`json:"bed_type"`
	Price		float64					`json:"price"`
}