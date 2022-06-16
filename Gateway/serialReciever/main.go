package main

import (
	"context"
	serialService "serialReciever/pkg/serialservice"
)

func main() {
	RecieverInit()
}

func RecieverInit() {
	serialConnection := serialService.ServiceServiceFactory()
	recievedData, recieveErr := serialConnection.Recieve()
	if recieveErr != nil {

	}
	ctx := context.Background()
	ctx, ctxCancel := context.WithCancel(ctx)

}
