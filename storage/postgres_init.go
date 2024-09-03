package storage

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func InitDB() *sqlx.DB {

	db, err := sqlx.Open("postgres", "")

	if err != nil {
		log.Fatal("Can't connect to DB!")
	}

	return db

}
