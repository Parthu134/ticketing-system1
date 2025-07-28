package controllers

import (
	"ticketing-system/config"
	"ticketing-system/models"
	"ticketing-system/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func VerifyOTP(c *fiber.Ctx) error {
	var payload struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}
	var user models.User
	if err := config.DB.Where("email=?", payload.Email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "invalid email",
		})
	}
	if user.OTP != payload.OTP || time.Now().After(user.OTPExpiry) {
		return c.Status(401).JSON(fiber.Map{
			"error": "invalid or expired OTP",
		})
	}
	user.OTP = ""
	user.OTPExpiry = time.Time{}
	config.DB.Save(&user)

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to generate token",
		})
	}
	return c.JSON(fiber.Map{
		"message": "login successful",
		"token":   token,
	})
}
