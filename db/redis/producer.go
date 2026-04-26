package db

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

func Init() (*Redis, error) {
	c := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	sc := c.Ping(context.Background())
	res := sc.String()
	_, err := fmt.Printf("Connected to redis succesed: %s\n", res)
	if err != nil {
		return nil, err
	}

	return &Redis{Client: c}, nil
}

func (redis *Redis) Producer(channel string, msg []byte) {
	ic := redis.Client.Publish(context.Background(), channel, msg)
	res := ic.String()
	_, err := fmt.Printf("Published [%s] : [%s]\n", channel, res)
	if err != nil {
		log.Fatal(err)
	}
}
