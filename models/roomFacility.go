package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoomFacility struct {
	RoomFacilityID	primitive.ObjectID		`bson:"_id,omitempty"`
	RoomID			primitive.ObjectID 		`bson:"_id,omitempty"`
	FacilityID		primitive.ObjectID 		`bson:"_id,omitempty"`
	Quantity		int						`json:"quantity"`
}