package tickets

import (
	"backend/models"
	"backend/postgres"
	"encoding/json"
	"net/http"
)

func CreatePoliceTicket(w http.ResponseWriter, r *http.Request) {
	var ticket models.Ticket
	json.NewDecoder(r.Body).Decode(&ticket)
	//save to the database
	db := postgres.DBConnect()
	db.Table("PoliceTickets").Create(&ticket)
	//we use this first ticket to add as the first update of the ticket
	var update models.UpdateTickets
	update.TicketID = ticket.TicketID
	update.UpdatedBy = "System"
	update.UpdateMsg = "Ticket Created"
	db.Table("TicketUpdates").Create(&update)
	json.NewEncoder(w).Encode(map[string]string{"status": "ticket created"})
}

func CreateAmbulanceTicket(w http.ResponseWriter, r *http.Request) {
	var ticket models.Ticket
	json.NewDecoder(r.Body).Decode(&ticket)
	//save to the database
	db := postgres.DBConnect()
	db.Table("AmbulanceTickets").Create(&ticket)
	//we use this first ticket to add as the first update of the ticket
	var update models.UpdateTickets
	update.TicketID = ticket.TicketID
	update.UpdatedBy = "System"
	update.UpdateMsg = "Ticket Created"
	db.Table("TicketUpdates").Create(&update)
	json.NewEncoder(w).Encode(map[string]string{"status": "ticket created"})
}

func CreateFireTicket(w http.ResponseWriter, r *http.Request) {
	var ticket models.Ticket
	json.NewDecoder(r.Body).Decode(&ticket)
	//save to the database
	db := postgres.DBConnect()
	db.Table("FireTickets").Create(&ticket)
	//we use this first ticket to add as the first update of the ticket
	var update models.UpdateTickets
	update.TicketID = ticket.TicketID
	update.UpdatedBy = "System"
	update.UpdateMsg = "Ticket Created"
	db.Table("TicketUpdates").Create(&update)
	json.NewEncoder(w).Encode(map[string]string{"status": "ticket created"})
}

func CreateUpdateTicket(w http.ResponseWriter, r *http.Request) {
	var ticket models.UpdateTickets
	json.NewDecoder(r.Body).Decode(&ticket)
	//save to the database
	db := postgres.DBConnect()
	db.Table("TicketUpdates").Create(&ticket)
	json.NewEncoder(w).Encode(map[string]string{"status": "ticket created"})
}
