package main

import (
	"backend/auth"
	"backend/controllers"
	"backend/tickets"
	"backend/websockets"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	//handling CORS issue since backend and frontend are running on different ports
	// credentials := handlers.AllowCredentials()
	// origins := handlers.AllowedOrigins([]string{"www.rescueai.amazonaws.com", "localhost:3000"})

	//auth routes
	r.HandleFunc("/api/v1/health", controllers.HealthCheck).Methods("GET")
	r.HandleFunc("/api/v1/signup", auth.Signup).Methods("POST")
	r.HandleFunc("/api/v1/login", auth.Login).Methods("POST")
	r.HandleFunc("/api/v1/logout", auth.Logout).Methods("GET")
	r.HandleFunc("/api/v1/check-token-expiry", auth.CheckExpiry).Methods("GET")
	r.HandleFunc("/api/v1/extend-token-expiry", auth.ExtendExpiry).Methods("GET")
	//route to get MY current http cookie
	r.HandleFunc("/api/v1/getmycookie", controllers.GetMyCookie).Methods("GET")

	//create table routes
	r.HandleFunc("/api/v1/createPoliceDB", tickets.CreatePoliceTable).Methods("GET")
	r.HandleFunc("/api/v1/createAmbulanceDB", tickets.CreateAmbulanceTable).Methods("GET")
	r.HandleFunc("/api/v1/createFireDB", tickets.CreateFireDepartmentTable).Methods("GET")
	r.HandleFunc("/api/v1/createUpdateTicket", tickets.TicketUpdateTables).Methods("GET")

	//create ticket routes
	r.HandleFunc("/api/v1/createPoliceTicket", tickets.CreatePoliceTicket).Methods("POST")
	r.HandleFunc("/api/v1/createAmbulanceTicket", tickets.CreateAmbulanceTicket).Methods("POST")
	r.HandleFunc("/api/v1/createFireTicket", tickets.CreateFireTicket).Methods("POST")
	//add the ticket updates
	r.HandleFunc("api/v1/ticketUpdate", tickets.CreateUpdateTicket).Methods("POST")

	//get all the tickets
	r.HandleFunc("/api/v1/getPoliceTickets", tickets.GetPoliceTickets).Methods("GET")
	r.HandleFunc("/api/v1/getAmbulanceTickets", tickets.GetAmbulanceTickets).Methods("GET")
	r.HandleFunc("/api/v1/getFireTickets", tickets.GetFireTickets).Methods("GET")
	//get the ticket updates
	r.HandleFunc("/api/v1/getTicketUpdates", tickets.GetTicketUpdates).Methods("GET")

	//handling websockets
	newHub := websockets.NewHub()
	go newHub.Run()
	r.HandleFunc("/api/v1/createWS", func(w http.ResponseWriter, r *http.Request) {
		websockets.ServerWs(newHub, w, r)
	}).Methods("GET")

	//storing messages
	r.HandleFunc("/api/v1/createPoliceMessageTable", websockets.CreatePoliceTable).Methods("GET")

	//retrievingmessages
	r.HandleFunc("/api/v1/getPoliceMessages", websockets.RetrievePoliceMessages).Methods("GET")

	log.Println("Server Running on port 5001")
	log.Fatal(http.ListenAndServe(":5001", r))
	// log.Fatal(http.ListenAndServe(":5001", handlers.CORS(credentials,origins)(r)))

}
