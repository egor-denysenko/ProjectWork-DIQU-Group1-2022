package main

import (
	"context"
	"fmt"
	"log"
	"serialReciever/pkg/protocolParser"
	serialService "serialReciever/pkg/serialservice"
	"time"
)

func main() {
	serialReaderInstance := RecieverInit()

	serialDataChan := make(chan []byte, 10)
	ctx := context.Background()
	ctx, ctxCancel := context.WithCancel(ctx)
	serialReaderInstance.Recieve(ctx, serialDataChan)

	go protocolParser.ValidateSerialData(ctx, serialDataChan)

	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("done")
	case <-ctx.Done():
		log.Println("context done")
		ctxCancel()
		break
	case <-serialDataChan:
		log.Println("ricevuto serial")
		protocolParser.ValidateSerialData(ctx, serialDataChan)
		log.Println(err)
	}
}

func RecieverInit() *serialService.SerialService {
	return serialService.ServiceServiceFactory("COM11")
}
