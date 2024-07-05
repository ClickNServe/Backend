package controllers

import (
	"backend/database"
	"backend/errors"
	"backend/models"
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUserWishlist(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	email := c.Query("email")
	if email == "" {
		return errors.GetError(c, "Need email as parameter!")
	}

	fmt.Println(email)

	collection := database.GetDatabase().Collection("wishlists")

	filter := bson.M{"useremail": email}
	
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return errors.GetError(c, "Error while find user wishlists")
	}
	defer cursor.Close(ctx)

	var wishlists []models.Wishlist
	err = cursor.All(ctx, &wishlists)
	if err != nil {
		return errors.GetError(c, "Error while decode wishlists")
	}

	return c.Status(fiber.StatusOK).JSON(wishlists)
}