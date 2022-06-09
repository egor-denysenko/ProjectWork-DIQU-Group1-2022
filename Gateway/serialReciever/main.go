package main

import (
	serialService "serialReciever/pkg/serialservice"
)

func main() {
	RecieverInit()
}

func RecieverInit() {
	serialConnection := serialService.ServiceServiceFactory()
	serialConnection.Recieve()
	serialConnection.Close()
}
