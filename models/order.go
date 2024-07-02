package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	OrderID			primitive.ObjectID		`json:"id" bson:"_id,omitempty"`
	Rooms			[]Room					`json:"rooms" bson:"rooms,omitempty"`
	UserEmail		string					`json:"email"`
	CheckIn			time.Time				`json:"check_in"`
	CheckOut		time.Time				`json:"check_out"`
	OrderTime		time.Time				`json:"order_time"`
	TotalCharge		float64					`json:"total_charge"`
	IsApproved		int						`json:"is_approved"`
	IsCanceled		int						`json:"is_canceled"`
}