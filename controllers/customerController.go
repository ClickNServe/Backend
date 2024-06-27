package controllers

import (
	"backend/database"
	"backend/models"
	"context"
	"log"
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
		log.Fatalf("Error while parsing data order")
	}

	checkIn, err := time.Parse(time.RFC3339, order.CheckIn.Format(time.RFC3339))
	if err != nil {
		log.Fatalf("Error while parsing check in time")
	}

	checkOut, err := time.Parse(time.RFC3339, order.CheckOut.Format(time.RFC3339))
	if err != nil {
		log.Fatalf("Error while parsing check out time")
	}

	order.CheckIn = checkIn
	order.CheckOut = checkOut
	order.OrderTime = time.Now()
	order.IsApproved = false
	order.IsCanceled = false
		
	collection := database.ConnectDatabase("orders")
	defer database.DisconnectDatabase()

	result, err := collection.InsertOne(ctx, order)
	if err != nil {
		log.Fatalf("Error while insert order to database")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func CancelReservation(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatalf("Invalid format ID")
	}

	collection := database.ConnectDatabase("orders")
	defer database.DisconnectDatabase()

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"is_canceled": true}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatalf("Failed update reservation")
	}
	
	return c.Status(fiber.StatusOK).JSON(result)
}

func AddToWishlist(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.Status(fiber.StatusOK).SendString("Successfully add to wishlist room id : " + id)
}

func DropRoomWishlist(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.Status(fiber.StatusOK).SendString("Successfully drop room wishlist id : " + id)
}