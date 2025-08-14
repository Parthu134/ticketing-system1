package controllers

import (
	"fmt"
	"math/rand"
	"strings"
	"ticketing-system/config"
	"ticketing-system/models"
	"ticketing-system/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&credentials); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "invalid request body",
			"details": err.Error(),
		})
	}
	if credentials.Email == "" || credentials.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "email and password are required",
		})
	}
	var user models.User
	if err := config.DB.Where("email= ? And password=?", credentials.Email, credentials.Password).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid email or password",
		})
	}
	credentials.Email = strings.TrimSpace(strings.ToLower(credentials.Email))
	credentials.Password = strings.TrimSpace(credentials.Password)
	otp := fmt.Sprintf("%5d", rand.Intn(10000))
	expiry := time.Now().Add(5 * time.Minute)
	user.OTP = otp
	user.OTPExpiry = expiry
	config.DB.Save(&user)
	if err := utils.SendOtp(user.Email, otp); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to send otp",
		})
	} 
	return c.JSON(fiber.Map{
		"message": "OTP sent to your mail",
	})
}
