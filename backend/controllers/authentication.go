package controllers

import (
	"backend/models"
	"backend/postgres"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("my-secret-key")

func CreateToken(username string) string {
	newToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"expires":  time.Now().Add(time.Hour * 24).Unix(),
		})
	signedJWT, err := newToken.SignedString(secretKey)
	if err != nil {
		log.Fatal(err)
	}
	return signedJWT
}

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
				Expires:  time.Now().Add(time.Hour * 24),
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"status": "login successful"})
		}
	}
}
