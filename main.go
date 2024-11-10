package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// struct for mux router
type route struct {
	Route *mux.Router
}

// struct for items class or object, using generics 'I' with the value as comparable
type items[I comparable] struct {
	ID    I       `json:"ID"`
	Name  string  `json:"Name"`
	Price float64 `json:"Price"`
}

// A methode for the route struct, creates a new instance of the mux router and assign the mux.NewRouter value to it,
// also startes the server with the http.ListenAndServe method
func (route *route) handleRoutes() {
	route.Route = mux.NewRouter().StrictSlash(true)

	// the /items endpoint and function to eb called when the endpoint gets an hit
	route.Route.HandleFunc("/items", getItems).Methods("GET")

	// starting a server with the http package and making use of the mux router instead of the default router which value would be nil
	http.ListenAndServe("localhost:9090", route.Route)

}

// Methode for GET request to the /items endpoint
func getItems(res http.ResponseWriter, req *http.Request) {
	// slice of new instance of the items struct class or object
	items := []items[string]{{ID: "1", Name: "Buscuit", Price: 10.00}}

	// parses the slice of into a json and returns it on our web page as our response data, it returns an error if the data passed to it is not successfully parsed
	err := json.NewEncoder(res).Encode(items)

	// checking if the data is successfully parsed into json, otherwise return an error
	if err != nil {
		http.Error(res, "Unabke to format json string, some error occured", http.StatusInternalServerError)
		return
	}

	// log statement for successful request handling
	log.Println("Successful - Status 200")
}

func main() {

	// A new instance of the mux router which the handleRoutes() function set a value to the route key in thr struct
	router := route{}

	// using the above route instance to make handleRoutes() it method
	router.handleRoutes()

}
