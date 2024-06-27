package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Facility struct {
	FacilityID		primitive.ObjectID 		`json:"id" bson:"_id,omitempty"`
	FacilityName	string					`json:"facility_name"`
	Price			int						`json:"price"`
}