package controllers

import "github.com/gofiber/fiber/v2"

func RegisterUser(c *fiber.Ctx) error {
	return c.SendString("Register User")
}
