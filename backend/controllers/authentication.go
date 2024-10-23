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
		log.Println("User already exists")
		return
	}

	//check if password and confirm password match
	if user.Password != user.ConfirmPassword {
		log.Println("Password and Confirm Password do not match")
		return
	}

	//save to the database
	db.Create(&user)
	log.Println("User created successfully")

}
