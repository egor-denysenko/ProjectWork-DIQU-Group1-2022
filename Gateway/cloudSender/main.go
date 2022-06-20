package main

import (
	"cloudSender/pkg/queueservice"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	redisConnection := queueservice.FactoryQueueService()
	data := redisConnection.Dequeue(ctx, "test")
	fmt.Println(string(data))
}
