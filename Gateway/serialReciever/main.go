package main

import (
	"context"
	"log"
	"serialReciever/pkg/protocolParser"
	"serialReciever/pkg/queueservice"
	serialService "serialReciever/pkg/serialservice"
)

func main() {
	serialReaderInstance := RecieverInit()
	queueService := QueueInit()
	err := queueService.Connect()
	if err != nil {
		log.Println(err)
	}
	ctx := context.Background()
	serialDataChan := make(chan []byte)
	parsedDataChan := make(chan []byte)
	go serialReaderInstance.Recieve(ctx, serialDataChan)

	for {
		select {
		case <-ctx.Done():
			log.Println("context done")
			break
		case testSerial := <-serialDataChan: //serialDataValue := <-serialDataChan:
			log.Println("ricevuto seriiale???")
			go protocolParser.ValidateSerialData(testSerial, parsedDataChan)
		case parsedData := <-parsedDataChan:
			log.Println("dato parsato")
			go queueService.Enqueue(ctx, "test", parsedData)
		}
	}
}

func RecieverInit() *serialService.SerialService {
	return serialService.ServiceServiceFactory("/dev/Uart485Dongle")
}

func QueueInit() *queueservice.QueueService {
	return queueservice.QueueServiceFactory()
}
