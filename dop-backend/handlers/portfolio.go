// portfolio.go
package handlers

import (
	"encoding/json"
	"net/http"
	"dop-backend/database"  // Replace with the actual path to your database package
)

// HandlePortfolio fetches the portfolio content from the database and sends it as a response
func HandlePortfolio(w http.ResponseWriter, r *http.Request) {
	// Fetch the portfolio content from the database
	portfolio, err := database.GetPortfolioContent()
	if err != nil {
		// If there is an error fetching the portfolio, return a 500 Internal Server Error
		http.Error(w, "Error fetching portfolio", http.StatusInternalServerError)
		return
	}

	// Set the response header to indicate it's JSON data
	w.Header().Set("Content-Type", "application/json")

	// Encode the portfolio struct as JSON and write it to the response
	err = json.NewEncoder(w).Encode(portfolio)
	if err != nil {
		// Handle JSON encoding errors
		http.Error(w, "Error encoding portfolio data", http.StatusInternalServerError)
	}
}
