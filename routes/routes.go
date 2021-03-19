
package routes

import (
	"net/http"

	"github.com/gothello/go-encurtador/controllers"
)

func LoadRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/home", controllers.Home)
	mux.HandleFunc("/api/url", controllers.Urls)
	mux.HandleFunc("/api/", controllers.Redirect)

	return mux
}
