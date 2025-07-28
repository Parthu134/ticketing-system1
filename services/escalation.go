package services

import (
	"fmt"
	"log"
	"ticketing-system/config"
	"ticketing-system/models"
	"time"
)

func AutoUpdatingStatus() {
	tickettime := time.Now().AddDate(-10, 0, 0)
	result := config.DB.Model(&models.Ticket{}).
		Where("status=? and given ticket <= ?", "open", tickettime).
		Update("status", "needs-attention")
	if result.Error != nil {
		log.Fatal("error", result.Error)
	} else if result.RowsAffected > 0 {
		log.Printf("%d tickets have been update to needs-attention", result.RowsAffected)
	}
}

func Updatetickets() {
	var tickets []models.Ticket

	timecut := time.Now().Add(-10 * time.Second)
	if err := config.DB.Where("status=? and created at<= ?", "open", timecut).Find(&tickets).Error; err != nil {
		log.Fatal("Error failed to fetch tickets", err)
	}
	for _, ticket := range tickets {
		ticket.Status = "need-attention"
		config.DB.Save(&ticket)

		notifications := models.Notifications{
			Role:     "admin",
			Title:    "ticket status updated to needs-attention",
			Message:  fmt.Sprintf("ticket %d titled %s needs attention", ticket.ID, ticket.Title),
			TicketID: ticket.ID,
		}
		config.DB.Create(&notifications)

		log.Printf("updated ticket %d and notified admin", ticket.ID)
	}
}
