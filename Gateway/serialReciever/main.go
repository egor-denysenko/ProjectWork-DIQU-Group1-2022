package main

import "serialReciever/pkg/serialReciever"

func main() {
	RecieverInit()
}

func RecieverInit() {

	serialReciever.NewSerialReciever()
}
