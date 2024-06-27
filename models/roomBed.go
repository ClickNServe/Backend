package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoomBed struct {
	RoomBedID	primitive.ObjectID		`json:"id" bson:"_id,omitempty"`
	RoomID		primitive.ObjectID 		`json:"room_id" bson:"room_id,omitempty"`
	BedID		primitive.ObjectID 		`json:"bed_id" bson:"bed_id,omitempty"`
	Quantity	int						`json:"quantity"`
}