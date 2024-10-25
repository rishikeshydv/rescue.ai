package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func CreateToken(username string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secret := os.Getenv("SECRET_KEY")
	var secretKey = []byte(secret)
	newToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"expires":  time.Now().Add(time.Hour * 1).Unix(),
		})
	signedJWT, err := newToken.SignedString(secretKey)
	if err != nil {
		log.Fatal(err)
	}
	return signedJWT
}
