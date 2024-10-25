package auth

import (
	"backend/models"
	"backend/postgres"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var loginUser models.LoginUser
	var dbUser models.LoginUser
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		log.Fatal(err)
	}
	//retrieve from the database if such a user exist
	db := postgres.DBConnect()
	tx := db.Table("signup_users").First(&loginUser, "phone=?", loginUser.Phone)
	log.Println(tx.RowsAffected)
	if tx.RowsAffected == 0 {
		json.NewEncoder(w).Encode(map[string]string{"status": "user does not exist"})
		return
	} else {
		//scan the database transaction into the dbUser
		tx.Scan(&dbUser)
		//check if the phone and password of the loginUser matches the one in the database transaction
		if loginUser.Phone == dbUser.Phone && loginUser.Password != dbUser.Password {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"status": "incorrect password"})
			return
		} else {
			tokenString := CreateToken(loginUser.Phone)
			//create a cookie to be passed
			cookie := http.Cookie{
				Name:     "token",
				Value:    tokenString,
				Expires:  time.Now().Add(time.Hour * 1),
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"status": "login successful"})
		}
	}
}
