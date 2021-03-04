package main

import (
	"fmt"
	"net/http"

	"github.com/gothello/go-encurtador/config"
	"github.com/gothello/go-encurtador/routes"
)

func main() {

	conf, err := config.Load()
	if err != nil {
		panic(err)
	}

	port := conf.GetInt("api_port")
	if port == 0 {
		port = 3000
	}

	mux := routes.LoadRoutes()

	server := &http.Server{
		Handler: mux,
		Addr:    fmt.Sprintf(":%d", port),
	}

	fmt.Printf("Server running port:%d\n", port)

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
