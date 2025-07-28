package controllers

import (
	"ticketing-system/config"
	"ticketing-system/models"

	"github.com/gofiber/fiber/v2"
)

func CreateTicket(c *fiber.Ctx) error {
	var input struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Category    string   `json:"category"`
		Status      string   `json:"status"`
		Priority    string   `json:"priority"`
		UserID      uint     `json:"user_id"`
		Tags        []string `json:"tags"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "invalid input data",
			"details": err.Error(),
		})
	}
	if input.Title == "" || input.Description == "" || input.UserID == 0 || input.Priority == "" || input.Category == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "missing required field",
		})
	}

	var tags []*models.Tag
	for _, tagName := range input.Tags {
		var tag models.Tag
		if err := config.DB.Where("name=?", tagName).FirstOrCreate(&tag, models.Tag{Name: tagName}).Error; err == nil {
			tags = append(tags, &tag)
		}
	}
	ticket := models.Ticket{
		Title:       input.Title,
		Description: input.Description,
		Category:    input.Category,
		Status:      "open",
		Priority:    input.Priority,
		UserID:      input.UserID,
		Tags:        tags,
	}
	if err := config.DB.Create(&ticket).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "failed to create ticket",
			"details": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(ticket)
}
