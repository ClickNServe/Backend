package controllers

import (
	"backend/database"
	"backend/errors"
	"backend/models"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllOrder(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var orders []models.Order

	collection := database.GetDatabase().Collection("orders")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return errors.GetError(c, "Error while find ID")
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &orders)
	if err != nil {
		return errors.GetError(c, "Error while decoding data")
	}

	return c.Status(fiber.StatusOK).JSON(orders)
}