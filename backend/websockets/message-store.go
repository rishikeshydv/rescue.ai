// we will be using this module to store the messages to the database
package websockets

import (
	"backend/models"
	"backend/postgres"
	"encoding/json"
	"log"
	"net/http"
)

func CreatePoliceTable(w http.ResponseWriter, r *http.Request) {
	db := postgres.DBConnect()
	db.Table("PoliceMessages").AutoMigrate(&models.Message{})
	json.NewEncoder(w).Encode(map[string]string{"status": "table created successfully"})
}

func AddPoliceMessages(newMsg *models.Message) {
	db := postgres.DBConnect()
	tx := db.Table("PoliceMessages").Create(&newMsg)
	if tx.Error != nil {
		log.Fatal(tx.Error.Error())
	}
	log.Println("Message Added Successfully")
}
