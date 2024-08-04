package utils

import (
	"encoding/json"
	"net/http"
)

// ParseJSONBody decodes the JSON body from the HTTP request into the provided interface.
// It handles any decoding errors by responding with a 400 Bad Request status.
func ParseJSONBody(w http.ResponseWriter, r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}
