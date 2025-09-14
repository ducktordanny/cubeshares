package main

import (
	"database/sql"
	"log"

	"github.com/ducktordanny/cubeshares/backend/cmd/api"
	"github.com/ducktordanny/cubeshares/backend/configs"
	"github.com/ducktordanny/cubeshares/backend/db"
)

func main() {
	config := db.GetConnectionURL()
	db, err := db.NewPostgresStorage(config)
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)

	server := api.NewAPIServer(configs.Envs.Port, configs.Envs.ClientAppURL, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Connected successfully!")
}
