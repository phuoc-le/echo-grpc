package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

type Client struct {
	Client *redis.Client
}

var ctx = context.Background()

func InitRedisClient(redisUrl string, password string, db int) (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: password,
		DB:       db,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Cannot Initialize Re*redis.Clientdis Client", err)
	}
	fmt.Println("Redis Client Successfully Initialized . . .", pong)

	return &Client{
		Client: client,
	}, nil
}
