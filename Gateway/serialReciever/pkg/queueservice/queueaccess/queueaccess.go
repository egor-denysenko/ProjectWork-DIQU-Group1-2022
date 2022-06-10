package queueaccess

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type IVagonMessage interface {
	Enqueue(context.Context, string, []byte)
}

type VagonMessageQueue struct {
	queueConnection IVagonMessage
}

func NewMessageQueue() *VagonMessageQueue {
	return &VagonMessageQueue{
		queueConnection: Connect(),
	}
}

func (v *VagonMessageQueue) Enqueue(ctx context.Context, message []byte) error {
	return nil
}

func Connect() IVagonMessage {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()
	rdb.LPush(ctx, "formaggio", "42Test")
	valore := rdb.RPop(ctx, "formaggio")
	fmt.Println(valore.Val())
	return rdb
}
