package main

import (
	"context"
	"log"
	"serialReciever/pkg/protocolParser"
	serialService "serialReciever/pkg/serialservice"
)

func main() {
	serialReaderInstance := RecieverInit()

	serialDataChan := make(chan []byte)

	recievedData, recieveErr := serialReaderInstance.Recieve()
	if recieveErr != nil {
		log.Fatalln(recieveErr)
	}
	serialDataChan <- recievedData
	ctx := context.Background()
	ctx, ctxCancel := context.WithCancel(ctx)
	protocolParser.ValidateSerialData(ctx, serialDataChan)
}

func RecieverInit() *serialService.SerialService {
	return serialService.ServiceServiceFactory("COM7")
}
