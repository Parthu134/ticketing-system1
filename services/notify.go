package services

import (
	"log"
	"ticketing-system/config"
	"ticketing-system/controllers"
	"ticketing-system/models"
	"time"
)

func CheckOldTicketsAndNotify() {
	var tickets []models.Ticket
	threedaysago := time.Now().AddDate(0, 0,-3)
	if err := config.DB.Where("status=? and created_at<?", "open", threedaysago).Find(&tickets).Error; err != nil {
		log.Fatal("Error failed old tickets", err)
		return
	}
	if len(tickets) == 0 {
		return
	}
	controllers.SendNotificationsToAdmin("Ticket alert", "There are tickets older than 3 days")
}
