package httpRespond

import (
	"encoding/json"
	"net/http"
)

type Respond interface {
	RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{})
	RespondWithError(w http.ResponseWriter, statusCode int, message string)
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	RespondWithJSON(w, statusCode, map[string]string{"error": message})
}
