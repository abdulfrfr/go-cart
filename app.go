package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type router struct {
	Route *mux.Router
	DB    *sql.DB
}

func (route *router) handleRequests() {
	route.Route.HandleFunc("/items", route.getItems).Methods("GET")
}

func (route *router) initialization() {
	route.Route = mux.NewRouter().StrictSlash(true)

	route.handleRequests()

	var err error
	route.DB, err = sql.Open("postgres", "user=gotest dbname=gotest password=pass#123 sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database ready for connection")

	defer route.DB.Close()

	log.Fatal(http.ListenAndServe("localhost:9090", route.Route))
}
