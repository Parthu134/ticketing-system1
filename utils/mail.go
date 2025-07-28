package utils

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func SendOtp(to string, otp string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "bheemanathiparthu08@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Your Login OTP")
	m.SetBody("text/plain", fmt.Sprintf("Your OTP is %s", otp))

	d := gomail.NewDialer("smtp.gmail.com", 587, "bheemanathiparthu08@gmail.com", "dkzd zsli qiuj djon")
	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println("SendOTP error", err)
	}
	return err
}
