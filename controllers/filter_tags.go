package controllers

import (
	"ticketing-system/config"
	"ticketing-system/models"

	"github.com/gofiber/fiber/v2"
)

func GetTicketsByTag(c *fiber.Ctx) error {
	tagName := c.Params("tag")
	var tag models.Tag
	if err := config.DB.
		Where("name=?", tagName).First(&tag).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "tag not found",
		})
	}
	var tickets []models.Ticket
	err := config.DB.
		Joins("JOIN ticket_tags ON ticket_tags.ticket_id = tickets.id").
		Where("ticket_tags.tag_id = ?", tag.ID).
		Preload("Tags").
		Find(&tickets).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not fetch tickets",
		})
	}

	return c.JSON(fiber.Map{
		"tag":     tag.Name,
		"tickets": tickets,
	})
}
