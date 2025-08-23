package controllers

import (
	"ticketing-system/config"
	"ticketing-system/models"

	"github.com/gofiber/fiber/v2"
)

func ListTicketsAdmins(c *fiber.Ctx) error {
	role := c.Locals("role", "admin")
	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "only admins can view all tickets",
		})
	}
	var tickets []models.Ticket
	if err := config.DB.Preload("Tags").Order("created_at desc").Find(&tickets).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "failed to fetch details",
			"details": err.Error(),
		})
	}
	return c.JSON(tickets)
}

func ReplyTicket(c *fiber.Ctx) error {
	paramsid := c.Params("id")
	var body struct {
		Response string `json:"response"`
	}
	if err := c.BodyParser(&body); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}
	var ticket models.Ticket
	if err := config.DB.First(&ticket, paramsid).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "connot find ticket",
		})
	}
	ticket.Response = body.Response
	ticket.Status = "closed"
	if err := config.DB.Save(&ticket).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not update the response",
		})
	}
	return c.JSON(fiber.Map{
		"message": "message from admin",
		"ticket":  ticket,

	})
}
