package redisClient

import (
	"log"

	"github.com/go-redis/redis/v8"
)

func Init(url string) *redis.Client{
	opt, err := redis.ParseURL(url)
	if err != nil {
		log.Fatal("Problem parsing redis url")
		panic(err)
	}
  	client := redis.NewClient(opt)
	return client
}