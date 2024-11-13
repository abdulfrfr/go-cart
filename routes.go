package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Items struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getItems(db *sql.DB) ([]Items, error) {

	rows, err := db.Query("SELECT * FROM list")

	if err != nil {
		return nil, err
	}

	itemRes := []Items{}

	for rows.Next() {
		eachItem := Items{}

		err = rows.Scan(&eachItem.ID, &eachItem.Name)

		if err != nil {
			return nil, err
		}

		itemRes = append(itemRes, eachItem)
	}

	return itemRes, nil

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
}
