package main

import (
	"cloudSender/pkg/mqttservice"
	"cloudSender/pkg/queueservice"
	"context"
	"log"
	"time"
)

func main() {
	ctx := context.Background()

	redisConnection := queueservice.FactoryQueueService()
	mqttService := mqttservice.FactoryMqttService()

	queueService := TryToConnectToQueue(redisConnection)
	mqttConnection := TryToConnectToBroker(mqttService)
	for {
		data, err := DequeueFromQueue(ctx, queueService)
		if err != nil {
			log.Println("error redis o dato nullo %v", err)
			continue
		}

		cloudSenderErr := mqttConnection.Pubblish("trainly/0/0/status", data, 0)
		if cloudSenderErr != nil {
			log.Println("Pubblish Unsuccessfull Retrying")
			enqueueErr := redisConnection.Enqueue(context.Background(), "test", []byte("test"))
			if enqueueErr != nil {
				log.Fatalln("can't enqueue not sent message")
			}
			log.Fatalln("connection unvalid with broker shutting down")
		}
	}
}

func TryToConnectToBroker(mqttService *mqttservice.MqttService) *mqttservice.MqttService {
	for {
		connectionErr := mqttService.Connect()
		if connectionErr != nil {
			time.After(1 * time.Second)
			continue
		}
		return mqttService
	}
}

func TryToConnectToQueue(redisConnection *queueservice.QueueService) *queueservice.QueueService {
	for {
		connectionErr := redisConnection.Connect()
		if connectionErr != nil {
			time.After(1 * time.Second)
			continue
		}
		return redisConnection
	}
}

func DequeueFromQueue(ctx context.Context, service *queueservice.QueueService) ([]byte, error) {
	return service.Dequeue(ctx, "test")
}
