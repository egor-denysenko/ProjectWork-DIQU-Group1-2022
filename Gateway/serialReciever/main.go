package main

import (
	"context"
	"log"
	serialService "serialReciever/pkg/serialservice"
)

func main() {
	serialReaderInstance := RecieverInit()
	ctx := context.Background()
	serialDataChan := make(chan []byte, 10)
	//parsedDataChan := make(chan []byte, 10)
	go serialReaderInstance.Recieve(ctx, serialDataChan)
	//go protocolParser.ValidateSerialData(serialDataChan, parsedDataChan)

	for {
		select {
		case <-ctx.Done():
			log.Println("context done")
			break
		case <-serialDataChan:
			//data := <-serialDataChan
			//protocolParser.ValidateSerialData(serialDataChan, parsedDataChan)
		}
	}
}

func RecieverInit() *serialService.SerialService {
	return serialService.ServiceServiceFactory("COM11")
}
