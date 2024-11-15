package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

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

func getItem(db *sql.DB, key string) (Items, error) {

	str := fmt.Sprintf("SELECT * FROM list WHERE id = %v", key)

	row, err := db.Query(str)

	if err != nil {
		log.Fatal(err)
	}

	var item Items
	for row.Next() {
		err = row.Scan(&item.ID, &item.Name)

		if err != nil {
			log.Fatal(err)
		}

	}
	return item, nil

}

func createItem(db *sql.DB, newItem *Items) ([]byte, error) {
	strCon := "INSERT INTO list (id, name) VALUES ($1, $2)"
	_, err := db.Exec(strCon, newItem.ID, newItem.Name)

	if err != nil {
		return nil, err
	}

	var res Items
	query := "SELECT id, name FROM list WHERE id = $1"
	err = db.QueryRow(query, newItem.ID).Scan(&res.ID, &res.Name)
	if err != nil {
		log.Fatalf("Failed to retrieve item: %v", err)
	}

	var result []byte
	var errs error
	result, errs = json.Marshal(res)

	if errs != nil {
		log.Fatal("Error parsing data")
	}

	return result, nil

}

func deleteItem(db *sql.DB, key string) (int64, error) {

	row, err := db.Exec("DELETE FROM list WHERE id = $1", key)

	if err != nil {
		log.Fatal(err)
	}
	var rowsAffected int64
	rowsAffected, err = row.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	return rowsAffected, nil

}

func updateItem(db *sql.DB, key string, item Items) error {

	_, err := db.Exec("UPDATE list SET name = $1 WHERE id = $2", item.Name, key)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil

}
