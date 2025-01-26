package api

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Message string `json:"message"`
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Create a sample JSON response
	response := ApiResponse{
		Message: "To jest API response w JSON!",
	}

	// Write the JSON response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
