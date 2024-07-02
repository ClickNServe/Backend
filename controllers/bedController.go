package controllers

import (
	"backend/database"
	"backend/errors"
	"backend/models"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllBed(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var beds []models.Bed

	collection := database.GetDatabase().Collection("beds")

	cursor, err := collection.Find(ctx, &beds)
	if err != nil {
		return errors.GetError(c, "Error while find ID")
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &beds)
	if err != nil {
		return errors.GetError(c, "Error while decoding data")	}

	return c.Status(fiber.StatusOK).JSON(beds)
}