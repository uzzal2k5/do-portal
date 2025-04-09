package routes

import (
	"github.com/gorilla/mux"
	"dop-backend/handlers"
)

// DopRouter initializes the router and routes
func DopRouter() *mux.Router {
	router := mux.NewRouter()

	// Services API
	router.HandleFunc("/api/services", handlers.GetAllServices).Methods("GET")
	router.HandleFunc("/api/services/{id}", handlers.GetServiceByID).Methods("GET")

	// Add more routes like:
	router.HandleFunc("/api/contact", handlers.HandleContactForm).Methods("POST")
	router.HandleFunc("/api/portfolio", handlers.HandlePortfolio).Methods("GET")

	return router
}

