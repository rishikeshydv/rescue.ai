package auth

import (
	"encoding/json"
	"net/http"
	"time"
)

func ExtendExpiry(w http.ResponseWriter, r *http.Request) {
	currentCookie, err := r.Cookie("token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}
	tokenString := currentCookie.Value
	newCookie := http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 1),
		HttpOnly: true,
	}
	http.SetCookie(w, &newCookie)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "logout successful"})
}
