package queueaccess

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
)

type VagonMessageQueue struct {
	queueConnection *redis.Client
}

func NewMessageQueue() *VagonMessageQueue {
	return &VagonMessageQueue{
		queueConnection: nil,
	}
}

func (v *VagonMessageQueue) Enqueue(ctx context.Context, key string, message []byte) error {
	_, err := v.queueConnection.RPush(ctx, key, message).Result()
	if err != nil {
		return err
	}
	return nil
}

func (v *VagonMessageQueue) Connect() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("RedisAddr"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	fmt.Println(err)
	if err != nil {
		return err
	}
	v.queueConnection = rdb
	return nil
}
