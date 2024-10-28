package handler

import "github.com/gofiber/fiber/v2"

func Alive(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status":"success", "message":"server is up!"})
}