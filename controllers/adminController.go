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

func CreateNewFacility(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var facility models.Facility
	err := c.BodyParser(&facility)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error while parsing data"})
	}

	collection := database.GetDatabase().Collection("facilities")

	result, err := collection.InsertOne(ctx, facility)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error while insert new facility"})
	}

	return c.Status(fiber.StatusBadRequest).JSON(result)
}

func UpdateFacility(c *fiber.Ctx) error {

}

func DeleteFacility(c *fiber.Ctx) error {

}

func CreateNewBed(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var bed models.Bed
	err := c.BodyParser(&bed)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error while parsing data"})
	}

	collection := database.GetDatabase().Collection("beds")

	result, err := collection.InsertOne(ctx, bed)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error while insert new data"})
	}

	return c.Status(fiber.StatusBadRequest).JSON(result)

}

func UpdateBed(c *fiber.Ctx) error {

}

func DeleteBed(c *fiber.Ctx) error {

}

func CreateNewRoom(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var room models.Room
	err := c.BodyParser(&room)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Eror while parsing data"})
	}

	collection := database.GetDatabase().Collection("rooms")

	result, err := collection.InsertOne(ctx, room)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error while insert new room"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func UpdateRoom(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.Status(fiber.StatusOK).SendString("Successfully update room number " + id + "!")
}

func DeleteRoom(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid format ID"})
	}

	collection := database.GetDatabase().Collection("rooms")

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"is_deleted": true}}

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error while delete room"})
	}

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Successfully deleted room!"})
}

func ApproveCustomerReservation(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Successfully approve customer reservation!")
}

func RejectCustomerReservation(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Successfully reject customer reservation!")
}

func HandleCustomerCancelation(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Sucessfully handled customer cancelation")
}