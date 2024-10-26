package tickets

import (
	"backend/models"
	"backend/postgres"
	"encoding/json"
	"net/http"
)

func CreatePoliceTable(w http.ResponseWriter, r *http.Request) {
	//connect to the database
	db := postgres.DBConnect()
	db.Table("PoliceTickets").AutoMigrate(models.Ticket{})
	json.NewEncoder(w).Encode(map[string]string{"status": "Police Table Created"})
}

func CreateAmbulanceTable(w http.ResponseWriter, r *http.Request) {
	//connect to the database
	db := postgres.DBConnect()
	db.Table("AmbulanceTickets").AutoMigrate(models.Ticket{})
	json.NewEncoder(w).Encode(map[string]string{"status": "Ambulance Table Created"})
}

func CreateFireDepartmentTable(w http.ResponseWriter, r *http.Request) {
	//connect to the database
	db := postgres.DBConnect()
	db.Table("FireDeptTickets").AutoMigrate(models.Ticket{})
	json.NewEncoder(w).Encode(map[string]string{"status": "Fire Department Table Created"})
}

func TicketUpdateTables(w http.ResponseWriter, r *http.Request) {
	//connect to the database
	db := postgres.DBConnect()
	db.Table("TicketUpdates").AutoMigrate(models.UpdateTickets{})
	json.NewEncoder(w).Encode(map[string]string{"status": "Ticket Update Table Created"})
}
