package controllers

import "github.com/gofiber/fiber/v2"


func ReserveRoom(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Successfully reserve room!")
}

func CancelReservation(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Successfully cancel room reservation")
}

func AddToWishlist(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.Status(fiber.StatusOK).SendString("Successfully add to wishlist room id : " + id)
}