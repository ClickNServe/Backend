package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	OrderID			primitive.ObjectID		`json:"id" bson:"_id,omitempty"`
	RoomID			[]Room					`json:"room_id" bson:"room_id,omitempty"`
	UserEmail		string					`json:"email"`
	CheckIn			time.Time				`json:"check_in"`
	CheckOut		time.Time				`json:"check_out"`
	OrderTime		time.Time				`json:"order_time"`
	TotalCharge		float64					`json:"total_charge"`
	IsApproved		bool					`json:"is_approved"`
	IsCanceled		bool					`json:"is_canceled"`
}