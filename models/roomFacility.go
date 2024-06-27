package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoomFacility struct {
	RoomFacilityID	primitive.ObjectID		`json:"id" bson:"_id,omitempty"`
	RoomID			primitive.ObjectID 		`json:"room_id" bson:"room_id,omitempty"`
	FacilityID		primitive.ObjectID 		`json:"facility_id" bson:"facility_id,omitempty"`
	Quantity		int						`json:"quantity"`
}