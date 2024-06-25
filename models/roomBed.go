package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoomBed struct {
	RoomBedID	primitive.ObjectID		`bson:"_id,omitempty"`
	RoomID		primitive.ObjectID 		`bson:"_id,omitempty"`
	BedID		primitive.ObjectID 		`bson:"_id,omitempty"`
	Quantity	int						`json:"quantity"`
}