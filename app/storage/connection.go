package storage

import (
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	DB *sqlx.DB
}

func NewStorage() (*Storage, error) {
	con, err := connectDB()
	if err != nil {
		return nil, err
	}

	return &Storage{DB: con}, nil

}

func connectDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "port=8088 user=postgres password=222 dbname=http-prjct sslmode=disable")
	if err != nil {
		panic(err)
	}
	if db != nil {
		db.Ping()
	}

	return db, nil
}