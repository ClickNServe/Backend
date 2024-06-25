package controllers

import "github.com/gofiber/fiber/v2"


func GetAllRoom(c* fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Successfully get all room!")
}

func GetAllAvailableRoom(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Sucessfully get all available room!")
}