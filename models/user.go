package models

import "time"

type User struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	Role      string `json:"role"`
	OTP       string
	OTPExpiry time.Time
}
