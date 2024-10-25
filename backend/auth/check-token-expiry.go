//we will be using this module to check the authentication status of the user
//if the cookie has been expired, the user will be redirected to the login page in the frontend

package auth

import (
	"encoding/json"
	"net/http"
	"time"
)

func CheckExpiry(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}
	tokenString := cookie.Value
	claims, err := VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}
	expiryTime := time.Unix(int64((claims["expires"].(float64))), 0)
	if time.Now().After(expiryTime) {
		json.NewEncoder(w).Encode(map[string]string{"status": "expired"})
	} else {
		json.NewEncoder(w).Encode(map[string]string{"status": "not-expired"})
	}

}
