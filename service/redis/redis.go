package redis

import (
	"os"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func Init() {
	//Create redis client
	Client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})
	_, err := Client.Ping().Result()
	if err != nil {
		panic(err)
	}
}
