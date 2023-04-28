package storage

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
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

func RedisDB() (*redis.Client, context.Context) {
	var ctx = context.Background()
	rds := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return rds, ctx
}
