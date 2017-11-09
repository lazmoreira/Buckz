package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func getNewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     c.RedisEndpoint,
		Password: c.RedisPass,
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}
