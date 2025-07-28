package models

import "gorm.io/gorm"

type Notifications struct {
	gorm.Model
	Role     string `json:"role"`
	Title    string `json:"title"`
	Message  string `json:"message"`
	TicketID uint   `json:"ticket_id"`
	Status   string `json:"status"`
}
