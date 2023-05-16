package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

func ConnectDB() (*sqlx.DB, error) {
	time.Sleep(5*time.Second)
	db, err := sqlx.Connect("postgres", "host=postgres port=5432 user=postgres password=222 dbname=post-prjct sslmode=disable")
	if err != nil {
		fmt.Println(err)
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
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	return rds, ctx
}
