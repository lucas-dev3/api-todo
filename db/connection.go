package db

import (
	"api-postgres/configs"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // _ is used to import the package without using it
)

func Connect() (*sql.DB, error) {
	conf := configs.GetDBConfig()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping()

	return conn, err
}
