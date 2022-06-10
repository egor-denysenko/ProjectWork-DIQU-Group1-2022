package queueaccess

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type VagonMessageQueue struct {
	queueConnection *redis.Client
}

func NewMessageQueue() *VagonMessageQueue {
	return &VagonMessageQueue{
		queueConnection: Connect(),
	}
}

func (v *VagonMessageQueue) Enqueue(ctx context.Context, key string, message []byte) error {
	return nil
}

func Connect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}
