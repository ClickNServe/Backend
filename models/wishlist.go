package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Wishlist struct {
	WishListID		primitive.ObjectID		`bson:"_id,omitempty"`
	RoomID			primitive.ObjectID		`bson:"_id,omitempty"`
	UserEmail		string					`json:"email"`
}