package middleware

import (
	"strings"
	"ticketing-system/utils"

	"github.com/gofiber/fiber/v2"
)

func RequiredAuth(c *fiber.Ctx) error {
	header := c.Get("Authorization")
	if header == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing authorization header",
		})
	}
	tokenstring := strings.Split(header, "Bearer ")
	if len(tokenstring) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid authorization token format",
		})
	}
	userID, role, err := utils.ParseJWT(tokenstring[1])
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid or expired token",
		})
	}
	c.Locals("userID", uint(userID))
	c.Locals("role", role)
	return c.Next()
}

func AdminOnly(c *fiber.Ctx) error {
	role := c.Locals("role", "admin")
	roleVal, ok := role.(string)
	if !ok || roleVal != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "only admins can access",
		})
	}
	return c.Next()
}
