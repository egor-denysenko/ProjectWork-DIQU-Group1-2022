package main

import (
	"MqttReceiver/pkg/mqttservice"
	"MqttReceiver/pkg/queueservice/queueaccess"
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	mqttSubChannel := make(chan []byte)
	mqttInstance := mqttservice.FactoryMqttService(mqttSubChannel)
	queueInstance := queueaccess.NewMessageQueue()
	brokerConnection := mqttInstance.Connect()
	log.Println(brokerConnection)
	log.Printf("%T %+v", brokerConnection, brokerConnection)

	queueConnectionErr := queueInstance.Connect()
	log.Println(queueConnectionErr)
	subscribeErr := mqttInstance.Subscribe(ctx, "trainly/+/+/status")
	log.Printf("Subscribe Error %T %+v", subscribeErr, subscribeErr)
	for {
		select {
		case mqttData := <-mqttSubChannel:
			log.Printf("hai ricevuto %v", mqttData)
			queueInstance.Enqueue(ctx, "testCommand", mqttData)
		}
	}
}
