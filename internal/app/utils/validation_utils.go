package utils

import (
	"encoding/json"
	"net/http"
)

// HandleResponse handles the HTTP response by setting the content type to JSON,
// writing the status code, and encoding the data into the response body.
func HandleResponse(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
