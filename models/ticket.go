package models

import (
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	UserID      uint   `json:"user_id"`
	Response    string `json:"response"`
	Tags        []*Tag `gorm:"many2many:ticket_tags" json:"tags"`
}
