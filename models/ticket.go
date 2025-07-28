package models

import "time"

type Ticket struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	UserID      uint      `json:"user_id"`
	Response    string    `json:"response"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
	Tags        []*Tag    `gorm:"many2many:ticket_tags" json:"tags"`
}
