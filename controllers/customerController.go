package controllers

import (
	"backend/database"
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error while parsing order data"})
	}

	id := c.Params("id")
	roomID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id format"})
	}

	checkIn, err := time.Parse(time.RFC3339, order.CheckIn.Format(time.RFC3339))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error while parsing check in time"})
	}

	checkOut, err := time.Parse(time.RFC3339, order.CheckOut.Format(time.RFC3339))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error while parsing check out time"})
	}

	order.RoomID = append(order.RoomID, roomID)
	order.CheckIn = checkIn
	order.CheckOut = checkOut
	order.OrderTime = time.Now()
		
	collection := database.GetDatabase().Collection("orders")

	result, err := collection.InsertOne(ctx, order)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error while insert order to database"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func CancelReservation(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid format ID"})
	}

	collection := database.GetDatabase().Collection("orders")

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"is_canceled": true}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed update reservation status"})
	}
	
	return c.Status(fiber.StatusOK).JSON(result)
}

func AddToWishlist(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wishlist models.Wishlist
	err := c.BodyParser(&wishlist)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error while parsing data"})
	}

	collection := database.GetDatabase().Collection("wishlists")

	roomID, err := primitive.ObjectIDFromHex(wishlist.RoomID.Hex())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid room ID"})
	}
	wishlist.RoomID = roomID

	result, err := collection.InsertOne(ctx, wishlist)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error while insert new document"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func DropRoomWishlist(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid format ID"})
	}

	collection := database.GetDatabase().Collection("wishlists")

	filter := bson.M{"_id": objectID}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error while drop wishlist"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}