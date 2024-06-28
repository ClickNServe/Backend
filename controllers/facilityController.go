package controllers

import (
	"backend/database"
	"backend/models"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllFacilities(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var facilities []models.Facility

	collection := database.GetDatabase().Collection("facilities")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error while find ID"})
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &facilities)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error while decoding data"})
	}

	return c.Status(fiber.StatusOK).JSON(facilities)
}