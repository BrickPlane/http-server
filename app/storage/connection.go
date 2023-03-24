package storage

import (
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
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
	godotenv.Load("secret.env")
	db, err := sqlx.Connect("postgres", os.Getenv("CONNECT"))
	if err != nil {
		panic(err)
	}
	if db != nil {
		db.Ping()
	}
	// db.MustExec(postCred)

	return db, nil
}

// const postCred = `
// CREATE TABLE IF NOT EXISTS "user" (
// 	id SERIAL PRIMARY KEY,
// 	email text UNIQUE,
// 	login text UNIQUE,
// 	password text
// );`
