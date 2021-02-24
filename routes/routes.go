package routes

import (

	"net/http"
	"github.com/gothello/go-encurtador/controllers"
)

func LoadRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controllers.Home)
	mux.HandleFunc("/urls", controllers.Urls)

	return mux
}