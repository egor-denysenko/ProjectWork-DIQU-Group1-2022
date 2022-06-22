package main

import (
	"cloudSender/pkg/mqttservice"
	"cloudSender/pkg/queueservice"
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	redisConnection := queueservice.FactoryQueueService()
	cloudSender := mqttservice.FactoryMqttService()
	connectionErr := cloudSender.Connect()

	for {
		data := EnqueueFromQueue(ctx, redisConnection)

		if connectionErr != nil {
			log.Fatalln("Connection error to MQTT broker")
		}

		cloudSenderErr := cloudSender.Pubblish("trainly/0/0/status", data, 0)
		if cloudSenderErr != nil {
			log.Println("Pubblish Unsuccessfull Retrying")
			enqueueErr := redisConnection.Enqueue(context.Background(), "test", []byte("test"))
			if enqueueErr != nil {
				log.Fatalln("can't enqueue not sent message")
			}
		}

	}
}

func EnqueueFromQueue(ctx context.Context, service *queueservice.QueueService) []byte {
	return service.Dequeue(ctx, "test")
}
