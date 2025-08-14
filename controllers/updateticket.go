package controllers

import (
	"strconv"
	"ticketing-system/config"
	"ticketing-system/models"

	"github.com/gofiber/fiber/v2"
)

func UpdateTicketStatus(c *fiber.Ctx) error {
	ticketID := c.Params("id")
	id, err := strconv.Atoi(ticketID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid ticket ID",
		})
	}
	var body struct {
		Status string `json:"status"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}
	allowedstatus := map[string]bool{
		"open":     true,
		"resolved": true,
	}
	if !allowedstatus[body.Status] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid status",
		})
	}
	var ticket models.Ticket
	if err := config.DB.Preload("Tags").First(&ticket, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "ticket not found",
		})
	}

	userID := c.Locals("userID")
	roleVal := c.Locals("role")

	if userID == nil || roleVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized: missing user context",
		})
	}

	_, ok := userID.(uint)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "invalid userID type",
		})
	}

	_, ok = roleVal.(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "invalid role type",
		})
	}

	ticket.Status = body.Status
	if err := config.DB.Save(&ticket).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update ticket",
		})
	}

	return c.JSON(fiber.Map{
		"message": "ticket status updated successfully",
		"ticket":  ticket,
	})
}
