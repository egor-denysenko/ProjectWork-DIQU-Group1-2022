package main

import (
	"SerialSender/pkg/protocolParser"
	"SerialSender/pkg/queueservice"
	serialService "SerialSender/pkg/serialservice"
	"context"
	"log"
	"os"
	"time"
)

func main() {
	serialReaderInstance := RecieverInit()
	queueInstance := QueueInit()
	queueConnection := TryToConnectToQueue(queueInstance)
	ctx := context.Background()
	queueDataChan := make(chan *string)
	parsedDataChan := make(chan []byte)
	go queueConnection.Dequeue(ctx, "test", queueDataChan)
	for {
		select {
		case <-ctx.Done():
			log.Println("context done")
			break
		case jsonToParse := <-queueDataChan: //serialDataValue := <-serialDataChan:
			log.Println("ricevuto Byte da coda?")
			if jsonToParse != nil {
				log.Print("%T %s", *jsonToParse)
				go protocolParser.ParseMessageToByte(jsonToParse, parsedDataChan)
			}
		case parsedData := <-parsedDataChan:
			if parsedData != nil {
				log.Printf("dato parsato: %v", parsedData)
				go serialReaderInstance.Send(ctx, parsedData)
			}
		}
	}
}

func RecieverInit() *serialService.SerialService {
	return serialService.ServiceServiceFactory(os.Getenv("SerialPortToListen"))
}

func QueueInit() *queueservice.QueueService {
	return queueservice.FactoryQueueService()
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
