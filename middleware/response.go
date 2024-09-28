package middleware

import (
	"encoding/json"
	"net/http"
)
type errorResponse struct {
	Err string `json:"err"`
}

func respond(w http.ResponseWriter, respData interface{}, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if respData != nil {
		err := json.NewEncoder(w).Encode(respData)
		if err != nil {
			http.Error(w, "Could not encode in json", http.StatusBadRequest)
			return
		}
	}
}