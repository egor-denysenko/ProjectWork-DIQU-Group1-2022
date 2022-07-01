package main

import (
	"SerialSender/pkg/protocolParser"
	"SerialSender/pkg/queueservice"
	serialService "SerialSender/pkg/serialservice"
	"context"
	"log"
)

func main() {
	serialReaderInstance := RecieverInit()
	queueService := QueueInit()
	err := queueService.Connect()
	if err != nil {
		log.Println(err)
	}
	ctx := context.Background()
	queueDataChan := make(chan *string)
	parsedDataChan := make(chan []byte)
	go queueService.Dequeue(ctx, "test", queueDataChan)
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
			log.Printf("dato parsato: %v", parsedData)
			go serialReaderInstance.Send(ctx, parsedData)
		}
	}
}

func RecieverInit() *serialService.SerialService {
	return serialService.ServiceServiceFactory("COM10") //"/dev/Uart485Dongle"
}

func QueueInit() *queueservice.QueueService {
	return queueservice.FactoryQueueService()
}
