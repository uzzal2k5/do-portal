package handlers

import (
	"encoding/json"
	"net/http"
	"dop-backend/services"
	"github.com/gorilla/mux"
	"strconv"
)

// GetAllServicesHandler handles the /api/services GET request
func GetAllServices(w http.ResponseWriter, r *http.Request) {
	servicesData, err := services.GetAllServices()
	if err != nil {
		http.Error(w, "Failed to retrieve services: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(servicesData)
}


// GetServiceByIDHandler handles the /api/services/{id} GET request
func GetServiceByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid service ID", http.StatusBadRequest)
		return
	}

	service, err := services.GetServiceByID(id)
	if err != nil {
		http.Error(w, "Service not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service)
}
