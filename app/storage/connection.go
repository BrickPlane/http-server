package storage

import (
	"github.com/jmoiron/sqlx"
)

func ConnectDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "port=8088 user=postgres password=222 dbname=http-prjct sslmode=disable")
	if err != nil {
		panic(err)
	}
	if db != nil {
		db.Ping()
	}

	return db, nil
}