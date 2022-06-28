package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mys/go-rentals/src/routes"
)

func main() {
	r := mux.NewRouter()
	routes.CarsRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
