package main

import (
	"context"
	"log"
	"serialReciever/pkg/protocolParser"
	serialService "serialReciever/pkg/serialservice"
)

func main() {
	serialReaderInstance := RecieverInit()
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
		case test := <-parsedDataChan:
			log.Println("dato parsato")
			log.Println(test)
		}
	}
}

func RecieverInit() *serialService.SerialService {
	return serialService.ServiceServiceFactory("COM11")
}
