package websockets

import (
	"backend/models"
	"backend/postgres"
	"encoding/json"
	"net/http"
)

func RetrievePoliceMessages(w http.ResponseWriter, r *http.Request) {
	db := postgres.DBConnect()
	var messages []models.Message
	db.Table("PoliceMessages").Find(&messages)
	json.NewEncoder(w).Encode(messages)
}
