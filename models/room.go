package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Room struct {
	RoomID			primitive.ObjectID 		`json:"id" bson:"_id,omitempty"`
	RoomNumber		int 					`json:"room_number"`
	Description		string					`json:"description"`
	Floor			int						`json:"floor"`
	PricePerNight	float64					`json:"price_per_night"`
	Availability	bool					`json:"availability"`
	SizeArea		float64					`json:"size_area"`
}