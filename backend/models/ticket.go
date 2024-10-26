package models

import (
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	TicketID   string `gorm:"unique" json:"ticketID" bson:"ticketID"`
	ClientName string `json:"clientName" bson:"clientName"`
	Address    string `json:"address" bson:"address"`
	Phone      string `json:"phone" bson:"phone"`
	Nature     string `json:"nature" bson:"nature"`
	Summary    string `json:"summary" bson:"summary"`
}
