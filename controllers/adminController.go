package controllers

import "github.com/gofiber/fiber/v2"


func CreateNewRoom(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Successfully create new room!")
}

func UpdateRoom(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.Status(fiber.StatusOK).SendString("Successfully update room number " + id + "!")
}

func DeleteRoom(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.Status(fiber.StatusOK).SendString("Sucessfully delete room number " + id + "!")
}

func ApproveCustomerReservation(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Successfully approve customer reservation!")
}

func RejectCustomerReservation(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Successfully reject customer reservation!")
}