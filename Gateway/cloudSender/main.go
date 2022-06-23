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
		data, err := DequeueFromQueue(ctx, redisConnection)
		if err != nil {
			log.Fatalf("error redis o dato nullo %v", err)
		}

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

func DequeueFromQueue(ctx context.Context, service *queueservice.QueueService) ([]byte, error) {
	return service.Dequeue(ctx, "test")
}
