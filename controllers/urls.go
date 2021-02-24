package controllers

import (

	"net/http"
	"github.com/gothello/go-encurtador/utils"
)

func Urls(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		utils.ToJson(w, "Hello World", http.StatusOK)
	}
}