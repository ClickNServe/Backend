package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Room struct {
	RoomID			primitive.ObjectID 		`json:"id" bson:"_id,omitempty"`
	BedID 			[]primitive.ObjectID	`json:"bedId" bson:"bedId,omitempty"`
	FacilityID 		[]primitive.ObjectID	`json:"facilityId" bson:"facilityId,omitempty"`
	Picture			string					`json:"picture"`
	RoomNumber		int 					`json:"roomnumber"`
	Description		string					`json:"description"`
	Floor			int						`json:"floor"`
	PricePerNight	float64					`json:"pricepernight"`
	Availability	bool					`json:"availability"`
	SizeArea		float64					`json:"sizearea"`
}