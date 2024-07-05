package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	OrderID			primitive.ObjectID		`json:"id" bson:"_id,omitempty"`
	Rooms			[]Room					`json:"rooms" bson:"rooms,omitempty"`
	UserEmail		string					`json:"email"`
	CheckIn			time.Time				`json:"checkIn"`
	CheckOut		time.Time				`json:"checkOut"`
	OrderTime		time.Time				`json:"orderTime"`
	TotalCharge		float64					`json:"totalCharge"`
	IsApproved		int						`json:"isApproved"`
	IsCanceled		int						`json:"isCanceled"`
}