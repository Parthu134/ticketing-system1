package controllers

import (
	"ticketing-system/config"
	"ticketing-system/models"

	"github.com/gofiber/fiber/v2"
)

func GetAdminNotifications(c *fiber.Ctx) error {
	var notify []models.Notifications
	if err:=config.DB.Where("role=?", "admin").Order("created_at desc").Find(&notify).Error; err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":"failed to fetch notifications",
		})
	}
	return c.JSON(notify)
}