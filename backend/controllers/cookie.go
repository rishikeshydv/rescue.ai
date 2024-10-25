package controllers

import (
	"encoding/json"
	"net/http"
)

func GetMyCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	//r.Header.Get("Authorization") this would give the same thing as above
	//log.Println(r.Header)
	//log.Println(cookie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"cookie": cookie.Value})

}
