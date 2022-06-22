package main

import (
	"cloudSender/pkg/mqttservice"
	"cloudSender/pkg/queueservice"
	"context"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()
	redisConnection := queueservice.FactoryQueueService()
	data := redisConnection.Dequeue(ctx, "test")
	fmt.Println(string(data))
	cloudSender := mqttservice.FactoryMqttService()
	connectionErr := cloudSender.Connect()
	if connectionErr != nil {
		log.Fatalln("Connection error to MQTT broker")
	}
	cloudSender.Pubblish("message", "test", 0)
	if connectionErr != nil {
		log.Println("Pubblish Unsuccessfull Retrying")
	}
}
