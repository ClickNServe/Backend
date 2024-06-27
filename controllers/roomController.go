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


func GetAllRoom(c* fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var rooms []models.Room

	collection := database.ConnectDatabase("rooms")
	defer database.DisconnectDatabase()

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatalf("Error while find id: %v", err)
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &rooms)
	if err != nil {
		log.Fatalf("Error while decode room: %v", err)
	}
	return c.Status(fiber.StatusOK).JSON(rooms)
}

func GetAllAvailableRoom(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := database.ConnectDatabase("rooms")
	defer database.DisconnectDatabase()

	filter := bson.M{"availability":true}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatalf("Erorr while filtering room %s", err)
	}
	defer cursor.Close(ctx)

	var rooms []models.Room
	err = cursor.All(ctx, &rooms)
	if err != nil {
		log.Fatalf("Error while decode available room data")
	}

	return c.Status(fiber.StatusOK).JSON(rooms)
}

func GetRoomDetail(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatalf("Error invalid id format")
	}

	var room models.Room

	collection := database.ConnectDatabase("rooms")
	defer database.DisconnectDatabase()

	err = collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&room)
	if err != nil {
		log.Fatalf("Error while decode object room: %s", err)
	}

	return c.Status(fiber.StatusOK).JSON(room)
}