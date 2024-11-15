package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Items struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (route *router) getItems(res http.ResponseWriter, req *http.Request) {
	items, err := getItems(route.DB)

	if err != nil {
		log.Fatal(err)
	}

	var item []byte
	item, err = json.Marshal(items)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(item)

	log.Println("Successful - ", http.StatusOK)
}

func (route *router) getItem(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-type", "application/json")
	key := mux.Vars(req)["id"]

	var it []byte
	item, err := getItem(route.DB, key)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}

	it, err = json.Marshal(item)

	if err != nil {
		log.Fatal(err)
	}
	res.WriteHeader(http.StatusOK)
	res.Write(it)
	log.Println("Sucessful - ", http.StatusOK)
}

func (route *router) createItem(res http.ResponseWriter, req *http.Request) {
	var newItem Items
	err := json.NewDecoder(req.Body).Decode(&newItem)

	if err != nil {
		log.Fatal(err)
	}

	result, err := createItem(route.DB, &newItem)

	res.Header().Set("Content-type", "application/json")

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(result)

}

func (route *router) deleteItem(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-type", "application/json")

	key := mux.Vars(req)["id"]

	_, err := deleteItem(route.DB, key)

	if err != nil {
		log.Fatal(err)
	}

	msg := fmt.Sprintf("Item %v", key)

	result := map[string]string{
		"deleted": msg,
	}

	var item []byte

	item, err = json.Marshal(result)

	if err != nil {
		log.Fatal(err)
	}

	res.WriteHeader(http.StatusOK)

	res.Write(item)

	log.Println("deletion successful")
}

func (route *router) updateItem(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	var item Items
	err := json.NewDecoder(req.Body).Decode(&item)

	if err != nil {
		log.Fatal(err)
	}

	key := mux.Vars(req)["id"]
	err = updateItem(route.DB, key, item)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	key = fmt.Sprintf("Sucessfully updated %v", key)
	response := map[string]string{
		"msg": key,
	}

	var result []byte
	result, err = json.Marshal(response)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Sucessful - ", http.StatusOK)
	res.WriteHeader(http.StatusOK)
	res.Write(result)
}
