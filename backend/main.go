package backend

import (
	"backend/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("api/v1/health", controllers.HealthCheck).Methods("GET")
}
