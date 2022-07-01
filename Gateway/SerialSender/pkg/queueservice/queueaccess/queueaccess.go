package queueaccess

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type VagonMessageQueue struct {
	queueConnection *redis.Client
}

func NewMessageQueue() *VagonMessageQueue {
	return &VagonMessageQueue{
		queueConnection: nil, //Connect(),
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
		Addr:     "localhost:6379",
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

func (v *VagonMessageQueue) Dequeue(ctx context.Context, key string, queueDataChan chan<- *string) {
	for {
		data, err := v.queueConnection.BRPop(ctx, 30*time.Second, key).Result()
		log.Printf("Data: %v , Err: %v", data, err)
		if err != nil {
			queueDataChan <- nil
		}
		if len(data) > 0 {
			queueDataChan <- &data[1]
		}
		continue
	}
}
