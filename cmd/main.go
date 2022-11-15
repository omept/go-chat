package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ong-gtp/go-chat/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterAuthRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9010", r))

}
