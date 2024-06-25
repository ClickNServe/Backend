package controllers

import (
	"github.com/gofiber/fiber/v2"
)


func GetAllRoom(c* fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Successfully get all room!")
}

func GetAllAvailableRoom(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Sucessfully get all available room!")
}

func GetRoomDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.Status(fiber.StatusOK).SendString("Successfully get room detail for id " + id)
}