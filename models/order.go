package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	OrderID			primitive.ObjectID		`bson:"_id,omitempty"`
	RoomID			primitive.ObjectID		`bson:"_id,omitempty"`
	UserEmail		string					`json:"email"`
	CheckIn			time.Time				`json:"check_in"`
	CheckOut		time.Time				`json:"check_out"`
	OrderTime		time.Time				`json:"order_time"`
	TotalCharge		float64					`json:"total_charge"`
	IsApprove		bool					`json:"is_approve"`
}