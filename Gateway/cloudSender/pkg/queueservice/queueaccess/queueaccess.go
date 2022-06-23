package queueaccess

import (
	"context"
	"time"

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
	_, err := v.queueConnection.RPush(ctx, key, message).Result()
	if err != nil {
		return err
	}
	return nil
}

func Connect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

func (v *VagonMessageQueue) Dequeue(ctx context.Context, key string) ([]byte, error) {
	data, err := v.queueConnection.BRPop(ctx, 30*time.Second, key).Result()
	if err != nil {
		return nil, err
	}
	return []byte(data[1]), nil
}
