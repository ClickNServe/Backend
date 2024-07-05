package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Wishlist struct {
	WishListID		primitive.ObjectID		`json:"id" bson:"_id,omitempty"`
	RoomID			primitive.ObjectID		`json:"roomId" bson:"roomId,omitempty"`
	UserEmail		string					`json:"useremail"`
}