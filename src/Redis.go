package main

import (
	"gopkg.in/redis.v3"
)

func getRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:    "db:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}