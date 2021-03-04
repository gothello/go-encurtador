package routes

import (
	"net/http"

	"github.com/gothello/go-encurtador/controllers"
)

func LoadRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/", controllers.Home)
	mux.HandleFunc("/api/urls", controllers.Urls)

	return mux
}
