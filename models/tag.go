package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model 
	Name string `json:"name"`
	Tickets []*Ticket `gorm:"many2many:ticket_tags" json:"-"`
}