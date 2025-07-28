package controllers

import (
	"strconv"
	"ticketing-system/config"
	"ticketing-system/models"

	"github.com/gofiber/fiber/v2"
)

func DeleteTicket(c *fiber.Ctx) error {
	ticketIDStr := c.Params("id")
	ticketID, err := strconv.Atoi(ticketIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid ticket ID",
		})
	}

	var ticket models.Ticket
	if err := config.DB.First(&ticket, ticketID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "ticket not found",
		})
	}

	if err := config.DB.Delete(&ticket).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete ticket",
		})
	}

	return c.JSON(fiber.Map{
		"message": "ticket deleted successfully",
	})
}
