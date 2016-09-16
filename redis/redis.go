package redis

import (
	"errors"

	"gopkg.in/redis.v4"
)

var client *redis.Client

func Client() *redis.Client {
	return client
}

func NewClient() error {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	if client == nil {
		return errors.New("cannot connect to redis")
	}
	return nil
}
