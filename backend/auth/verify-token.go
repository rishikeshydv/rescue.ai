package auth

import (
	"log"
	"os"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func VerifyToken(tokenString string) (jwt.MapClaims,error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secretKey := os.Getenv("SECRET_KEY")
	//parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		log.Fatal(err)
	}
	if token.Valid {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Fatal(err)
		}
		return claims,nil
		//expiryTime := time.Unix(int64((claims["expires"].(float64))), 0)
	}
	return nil,err
}
