package db

import (
	"database/sql"
	"log"
)

func NewMySQLStorage(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
