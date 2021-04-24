package utils

import (
	"encoding/json"
	"net/http"
)

func ToJson(w http.ResponseWriter, message interface{}, status int) {

	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(status)

	json.NewEncoder(w).Encode(message)
}
