package messageEnqueuer

import (
	"context"
	"fmt"
	"serialReciever/internal/queueAbstraction"

	"github.com/go-redis/redis/v8"
)

type VagonMessage struct {
	ParsedMeggage []byte
}

type VagonMessageQueue struct {
	queue queueAbstraction.QueueActions
}

func NewMessageQueue(queue queueAbstraction.QueueActions) *VagonMessageQueue {
	return &VagonMessageQueue{}
}

func (v *VagonMessageQueue) Enqueue(VagonMessage) error {
	return nil
}

func NewQueue() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()
	rdb.LPush(ctx, "formaggio", "42Test")
	valore := rdb.RPop(ctx, "formaggio")
	fmt.Println(valore.Val())
}
