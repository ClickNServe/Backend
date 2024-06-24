package main

import (
	"backend/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SetUp(app)

	log.Fatal(app.Listen(":8080"))
	
}