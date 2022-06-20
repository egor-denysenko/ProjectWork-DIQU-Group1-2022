package mqttaccess

import mqtt "github.com/eclipse/paho.mqtt.golang"

type MqttClient struct {
	mqttConnection *mqtt.Client
}

func NewMqttConnection() *mqtt.Client {
	return mq
}
