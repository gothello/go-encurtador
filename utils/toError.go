package utils

import (

	"net/http"
)
func ToError(w http.ResponseWriter, message string, code int){
	ToJson(w, map[string]string{"error": message}, code)
}