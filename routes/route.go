package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	app.Get("/oauth/google", controllers.GoogleAuthEndpoint)
	app.Get("/oauth/redirect", controllers.GoogleRedirectEndpoint)
}
