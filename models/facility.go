package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Facility struct {
	FacilityID		primitive.ObjectID 		`json:"id" bson:"_id,omitempty"`
	FacilityName	string					`json:"facilityname"`
	Price			float64					`json:"price"`
}