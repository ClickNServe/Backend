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

func CreateNewFacility(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var facility models.Facility
	err := c.BodyParser(&facility)
	if err != nil {
		return errors.GetError(c, "Error while parsing data")
	}

	collection := database.GetDatabase().Collection("facilities")

	result, err := collection.InsertOne(ctx, facility)
	if err != nil {
		return errors.GetError(c, "Error while insert new data")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func UpdateFacility(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.GetError(c, "Invalid format ID")
	}

	var newFacility models.Facility
	err = c.BodyParser(&newFacility)
	if err != nil {
		return errors.GetError(c, "Error while parsing data")
	}

	collection := database.GetDatabase().Collection("facilities")
	filter := bson.M{"_id": objectId}
	update := bson.M{
		"$set": bson.M{
			"facility_name": newFacility.FacilityName,
			"price": newFacility.Price,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.GetError(c, "Error while updating data")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func DeleteFacility(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.GetError(c, "Invalid format ID")
	}

	collection := database.GetDatabase().Collection("facilities")
	filter := bson.M{"_id": objectId}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.GetError(c, "Error while deleting data")
	}

	collection = database.GetDatabase().Collection("rooms")
	filter = bson.M{"facility_id": objectId}
	update := bson.M{"$pull": bson.M{"facility_id": objectId}}

	_, err = collection.UpdateMany(ctx, filter, update)
	if err != nil {
		return errors.GetError(c, "Error while updating rooms")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func CreateNewBed(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var bed models.Bed
	err := c.BodyParser(&bed)
	if err != nil {
		return errors.GetError(c, "Error while parsing data")
	}

	collection := database.GetDatabase().Collection("beds")

	result, err := collection.InsertOne(ctx, bed)
	if err != nil {
		return errors.GetError(c, "Error while insert new data")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func UpdateBed(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.GetError(c, "Invalid format ID")
	}

	var bed models.Bed
	err = c.BodyParser(&bed)
	if err != nil {
		return errors.GetError(c, "Error while parsing data")
	}

	collection := database.GetDatabase().Collection("beds")
	filter := bson.M{"_id": objectId}
	update := bson.M{
		"$set": bson.M{
			"bed_type": bed.BedType,
			"price": bed.Price,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.GetError(c, "Error while updating data")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func DeleteBed(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.GetError(c, "Invalid format ID")
	}

	collection := database.GetDatabase().Collection("beds")

	filter := bson.M{"_id": objectId}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.GetError(c, "Error while deleting data")
	}

	collection = database.GetDatabase().Collection("rooms")
	filter = bson.M{"bed_id": objectId}
	update := bson.M{"$pull": bson.M{"bed_id": objectId}}

	_, err = collection.UpdateMany(ctx, filter, update)
	if err != nil {
		return errors.GetError(c, "Error while updating rooms")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func CreateNewRoom(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var room models.Room
	err := c.BodyParser(&room)
	if err != nil {
		return errors.GetError(c, "Error while parsing data")
	}

	collection := database.GetDatabase().Collection("rooms")

	result, err := collection.InsertOne(ctx, room)
	if err != nil {
		return errors.GetError(c, "Error while insert new data")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func UpdateRoom(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.GetError(c, "Error while parsing data")
	}

	var room models.Room
	err = c.BodyParser(&room)
	if err != nil {
		return errors.GetError(c, "Error while parsing data")
	}

	collection := database.GetDatabase().Collection("rooms")
	filter := bson.M{"_id": objectId}
	update := bson.M{
		"$set": bson.M{
			"bed_id": room.BedID,
			"facility_id": room.FacilityID,
			"picture": room.Picture,
			"room_number": room.RoomNumber,
			"description": room.Description,
			"floor": room.Floor,
			"price_per_night": room.PricePerNight,
			"availability": room.Availability,
			"size_area": room.SizeArea,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.GetError(c, "Error while updating data")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func DeleteRoom(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.GetError(c, "Error while parsing data")
	}

	collection := database.GetDatabase().Collection("rooms")
	filter := bson.M{"_id": objectId}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.GetError(c, "Error while deleting data")
	}

    return c.Status(fiber.StatusOK).JSON(result)
}

func ApproveCustomerReservation(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.GetError(c, "Invalid format ID")
	}

	collection := database.GetDatabase().Collection("orders")

	filter := bson.M{"_id": objectId}
	update := bson.M{
		"$set": bson.M{
			"is_approved": 1, 
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.GetError(c, "Error while updating data")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func RejectCustomerReservation(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.GetError(c, "Invalid format ID")
	}

	collection := database.GetDatabase().Collection("orders")
	filter := bson.M{"_id": objectId}
	update := bson.M{
		"$set": bson.M{
			"is_approved": 2,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.GetError(c, "Error while updating data")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}