package storage

import (
	"github.com/jmoiron/sqlx"
)

type DB struct {
	db *sqlx.DB
}

func InitDB() {

	db, err := sqlx.Open()

}
