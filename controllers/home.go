package controllers

import (

	"net/http"
//	"github.com/gothello/go-encurtador/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to API Shortener!"))
}