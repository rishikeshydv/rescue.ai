package models
import (
	"gorm.io/gorm"
)
type UpdateTickets struct {
	gorm.Model
	TicketID   string `gorm:"unique" json:"ticketID" bson:"ticketID"`
	UpdatedBy string `json:"updatedBy" bson:"updatedBy"`
	UpdateMsg string `json:"updateMsg" bson:"updateMsg"`
}