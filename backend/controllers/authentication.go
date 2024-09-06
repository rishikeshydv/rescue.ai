package controllers

import (
	"backend/models"
	"encoding/json"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var user models.SignupUser
	// Decode the incoming SignupUser json
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//check if the user already exists

	//check if password and confirm password match

	//save to the database

	//direct to the login page
}
