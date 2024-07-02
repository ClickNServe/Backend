package controllers

import (
	"backend/database"
	"backend/errors"
	"backend/models"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReserveRoom(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var order models.Order
	err := c.BodyParser(&order)
	if err != nil {
		return errors.GetError(c, "Error while parsing data")
	}

	checkIn, err := time.Parse(time.RFC3339, order.CheckIn.Format(time.RFC3339))
	if err != nil {
		return errors.GetError(c, "Error while parsing check in time")
	}

	checkOut, err := time.Parse(time.RFC3339, order.CheckOut.Format(time.RFC3339))
	if err != nil {
		return errors.GetError(c, "Error while parsing check out time")
	}

	order.CheckIn = checkIn
	order.CheckOut = checkOut
	order.OrderTime = time.Now()
		
	collection := database.GetDatabase().Collection("orders")

	result, err := collection.InsertOne(ctx, order)
	if err != nil {
		return errors.GetError(c, "Error while insert new data")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func AddToWishlist(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wishlist models.Wishlist
	err := c.BodyParser(&wishlist)
	if err != nil {
		return errors.GetError(c, "Error while parsing data")
	}
	
	roomID, err := primitive.ObjectIDFromHex(wishlist.RoomID.Hex())
	if err != nil {
		return errors.GetError(c, "Invalid format ID")
	}
	wishlist.RoomID = roomID

	collection := database.GetDatabase().Collection("wishlists")

	result, err := collection.InsertOne(ctx, wishlist)
	if err != nil {
		return errors.GetError(c, "Error while insert new data")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func DropRoomWishlist(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.GetError(c, "Invalid format ID")
	}

	collection := database.GetDatabase().Collection("wishlists")

	filter := bson.M{"_id": objectID}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.GetError(c, "Error while deleting data")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}