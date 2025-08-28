package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {

	reddisDb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("DB_ADDRESS"),
		Password: os.Getenv("DB_PASSWORD") ,
		DB: dbNo,
	})
 
	return reddisDb
}

