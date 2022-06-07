package main

import (
	"fmt"
	serialreciever "serialReciever/internal/serialReciever"
)

func main() {
	RecieverInit()
}

func RecieverInit() {
	//serialConfig := serialconfig.NewSerialConfig()

	serialPortConnection := serialreciever.NewSerialPortReader()

	rawSerialData, rawDataErr := serialPortConnection.RecieveSerial()

	if rawDataErr != nil {
		panic(rawDataErr)
	}
	fmt.Println(rawSerialData)
}
