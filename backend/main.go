package main

import (
	"backend/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/health", controllers.HealthCheck).Methods("GET")
	r.HandleFunc("/api/v1/signup", controllers.Signup).Methods("POST")
	r.HandleFunc("/api/v1/login", controllers.Login).Methods("POST")
	log.Println("Server Running on port 5001")
	log.Fatal(http.ListenAndServe(":5001", r))
}
