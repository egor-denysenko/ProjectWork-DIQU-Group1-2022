package main

import (
	"cloudSender/pkg/mqttservice"
	"cloudSender/pkg/queueservice"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	redisConnection := queueservice.FactoryQueueService()
	data := redisConnection.Dequeue(ctx, "test")
	fmt.Println(string(data))
	cloudSender := mqttservice.FactoryMqttService()
	cloudSender.Connect()
	cloudSender.Pubblish("message", "test", 0)
}
