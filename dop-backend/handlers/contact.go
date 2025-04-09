// contact.go
package handlers

import (
	"encoding/json"
	"net/http"
	"dop-backend/database"  // Replace with the actual path to your database package
	"dop-backend/models"    // Replace with the actual path to your models package
	"time"
)

// HandleContactForm handles POST requests for the contact form submission
func HandleContactForm(w http.ResponseWriter, r *http.Request) {
	var form models.ContactForm
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	form.CreatedAt = time.Now()
	if err := database.SaveContactForm(form); err != nil {
		http.Error(w, "Error saving contact form", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Contact form submitted successfully")
}

