package controllers

import (
	"ticketing-system/config"
	"ticketing-system/models"

	"github.com/gofiber/fiber/v2"
)

func SaveSubscription(c *fiber.Ctx) error {
	var sub struct {
		Endpoint string `json:"endpoint"`
		Keys     struct {
			P256dh string `json:"p256dh"`
			Auth   string `json:"auth"`
		} `json:"keys"`
	}
	if err := c.BodyParser(&sub); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}
	s := models.Subscription{
		Role:     "admin",
		Endpoint: sub.Endpoint,
		P256dh:   sub.Keys.P256dh,
		Auth:     sub.Keys.Auth,
	}
	if err := config.DB.Save(&s).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to save subscription",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "subscription saved",
	})
}
