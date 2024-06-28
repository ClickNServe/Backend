package main

import (
	"backend/database"
	"backend/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDatabase()
	defer database.DisconnectDatabase()

	routes.SetUp(app)

	log.Fatal(app.Listen(":8080"))
	
}