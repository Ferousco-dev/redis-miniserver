package main

import (
	"context"
	"log"
	"redis-miniserver/cache"
)

const (
	RedisAddr     string = "localhost:6379"
	RedisPassword string = ""
	RedisDB       int    = 0

)

func main() {
 ctx := context.Background()

 c := cache.New(RedisAddr, RedisPassword, RedisDB)
 if  err  := c.Ping(ctx); err != nil {
	log.Panic("failed to connect to Redis")
 }
 log.Println("connected to redis..")

 if err := c.Set(ctx, "user:name", "feranmi", 0); err != nil {
	log.Printf("Error: could not store a value in Redis")
  }

 log.Println("value stored in Redis")

 res, err := c.Get(ctx, "user:name")
	if err != nil {
		log.Println("Error: could  not get a free Redis")
	}

	log.Println("Result:", res)
}


