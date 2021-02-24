package utils

import (

	"net/http"
	"encoding/json"
)

func ToJson(w http.ResponseWriter, message interface{}, status int) {
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(message)
}