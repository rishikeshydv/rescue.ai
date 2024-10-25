package auth

import (
	"backend/models"
	"backend/postgres"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	//connect to the database
	db := postgres.DBConnect()
	db.AutoMigrate(models.SignupUser{})
	var user models.SignupUser
	// Decode the incoming SignupUser json
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(user)

	//check if the user already exists
	tx := db.First(&user, "phone=?", user.Phone)
	if tx.RowsAffected > 0 {
		json.NewEncoder(w).Encode(map[string]string{"status": "user already exists"})
		return
	}

	//check if password and confirm password match
	if user.Password != user.ConfirmPassword {
		json.NewEncoder(w).Encode(map[string]string{"status": "passwords do not match"})
		return
	}

	//save to the database
	tx = db.Create(&user)
	if tx.Error != nil {
		log.Fatal(tx.Error)
		return
	} else {
		json.NewEncoder(w).Encode(map[string]string{"status": "user created"})
	}

}
