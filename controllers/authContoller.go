package controllers

import (
	"backend/models"
	"backend/oauth"
	"context"
	"encoding/json"
	"io"

	"github.com/gofiber/fiber/v2"
)

func GoogleAuthEndpoint(c *fiber.Ctx) error {
	url := oauth.GoogleOAuthConfig().AuthCodeURL("state")
	return c.Redirect(url)
}

func GoogleRedirectEndpoint(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to exchange token")
	}

	token, err := oauth.GoogleOAuthConfig().Exchange(context.Background(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to exchange token: " + err.Error())
	}

	client := oauth.GoogleOAuthConfig().Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user info: " + err.Error())
	}

	defer response.Body.Close()

	var user models.User

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error reading response body: " + err.Error())
	}

	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error unmarshal json body: " + err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(user)
}