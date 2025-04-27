package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ducktordanny/cubeit/backend/configs"
)

func GetConnectionURL() string {
	name := configs.Envs.DBName
	user := configs.Envs.DBUser
	passwd := configs.Envs.DBPassword
	host := configs.Envs.DBHost
	port := configs.Envs.DBHostPort
	connectionURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, passwd, host, port, name)

	disableSSLMode := configs.Envs.DBDisableSSLMode
	if disableSSLMode == true {
		connectionURL = connectionURL + "?sslmode=disable"
	}
	return connectionURL
}

func NewPostgresStorage(config string) (*sql.DB, error) {
	db, err := sql.Open("postgres", GetConnectionURL())
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
