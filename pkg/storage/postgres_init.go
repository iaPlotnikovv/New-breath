package storage

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "test"
	dbname   = "mydb"
)

func InitDB() *sqlx.DB {

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sqlx.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Can't connect to DB!")
	}

	return db

}
