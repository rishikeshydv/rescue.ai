package main

import (
	"backend/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("api/v1/health", controllers.HealthCheck).Methods("GET")
	log.Println("Server Running on port 5001")
	log.Fatal(http.ListenAndServe(":5001", r))
}
