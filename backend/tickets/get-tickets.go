package tickets

import (
	"backend/models"
	"backend/postgres"
	"encoding/json"
	"net/http"
)

func GetPoliceTickets(w http.ResponseWriter, r *http.Request) {
	db := postgres.DBConnect()
	db.AutoMigrate(models.Ticket{})
	var tickets []models.Ticket
	tx := db.Table("PoliceTickets").Find(&tickets)
	if tx.Error != nil {
		json.NewEncoder(w).Encode(map[string]string{"status": "error"})
	}
	tx.Scan(&tickets)
	json.NewEncoder(w).Encode(tickets)
}

func GetAmbulanceTickets(w http.ResponseWriter, r *http.Request) {
	db := postgres.DBConnect()
	db.AutoMigrate(models.Ticket{})
	var tickets []models.Ticket
	tx := db.Table("AmbulanceTickets").Find(&tickets)
	if tx.Error != nil {
		json.NewEncoder(w).Encode(map[string]string{"status": "error"})
	}
	tx.Scan(&tickets)
	json.NewEncoder(w).Encode(tickets)
}

func GetFireTickets(w http.ResponseWriter, r *http.Request) {
	db := postgres.DBConnect()
	db.AutoMigrate(models.Ticket{})
	var tickets []models.Ticket
	tx := db.Table("FireDeptTickets").Find(&tickets)
	if tx.Error != nil {
		json.NewEncoder(w).Encode(map[string]string{"status": "error"})
	}
	tx.Scan(&tickets)
	json.NewEncoder(w).Encode(tickets)
}

func GetTicketUpdates(w http.ResponseWriter, r *http.Request) {
	var ticketID string
	err := json.NewDecoder(r.Body).Decode(&ticketID)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"status": "error"})
	}
	db := postgres.DBConnect()
	db.AutoMigrate(models.UpdateTickets{})
	var updates []models.UpdateTickets
	tx := db.Table("TicketUpdates").Where("ticketID=?", ticketID).Find(&updates).Order("updated_at ASC")
	if tx.Error != nil {
		json.NewEncoder(w).Encode(map[string]string{"status": "error"})
	}
	tx.Scan(&updates)
	json.NewEncoder(w).Encode(updates)
}
