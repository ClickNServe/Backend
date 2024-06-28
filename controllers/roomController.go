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


func GetAllRoom(c* fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var rooms []models.Room

	collection := database.GetDatabase().Collection("rooms")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error while find ID"})
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &rooms)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error while decode room"})
	}

	return c.Status(fiber.StatusOK).JSON(rooms)
}

func GetAllAvailableRoom(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := database.GetDatabase().Collection("rooms")

	filter := bson.M{"availability":true}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error while filtering room"})
	}
	defer cursor.Close(ctx)

	var rooms []models.Room
	err = cursor.All(ctx, &rooms)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error while decode room"})
	}

	return c.Status(fiber.StatusOK).JSON(rooms)
}

func GetRoomDetail(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid format id"})
	}

	var room models.Room

	collection := database.GetDatabase().Collection("rooms")

	err = collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&room)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error while decode room"})
	}

	return c.Status(fiber.StatusOK).JSON(room)
}