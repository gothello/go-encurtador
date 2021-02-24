package main

import (

	"fmt"
	"flag"
	"net/http"
	"github.com/gothello/go-encurtador/routes"
)

var (
	port = flag.Int("port", 3000, "port running service")
)

func init() {
	flag.Parse()
}

func main() {
	
	mux := routes.LoadRoutes()

	server := &http.Server{
		Handler: mux,
		Addr: fmt.Sprintf(":%d", *port),
	}

	fmt.Printf("Server running port:%d\n", *port)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
