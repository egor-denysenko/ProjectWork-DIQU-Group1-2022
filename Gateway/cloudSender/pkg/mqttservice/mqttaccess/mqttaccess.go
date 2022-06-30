package mqttaccess

import (
	"context"
	mqtt "github.com/eclipse/paho.golang/paho"
	"log"
	"net"
)

type MqttClient struct {
	mqttInstance   *mqtt.Client
	mqttConnection *mqtt.Connack
}

func NewMqttConnection() *MqttClient {
	serverUrl := "20.238.251.167:1883"
	// Try to reach the Broker and connects to it
	/*brokerDial, err := tls.Dial("tcp", serverUrl, &tls.Config{
		InsecureSkipVerify: true,
	})*/
	brokerDial, err := net.Dial("tcp", serverUrl)
	// manage the connection error from the Dial
	if err != nil {
		log.Fatalf("Failed to connect to %s: %s", serverUrl, err)
	}
	// Create a default broker instance
	brokerClient := mqtt.NewClient(mqtt.ClientConfig{
		Conn: brokerDial,
	})

	return &MqttClient{
		mqttInstance:   brokerClient,
		mqttConnection: nil,
	}
}

//topic := "/gne"
//qos := 2
func (m *MqttClient) Connect() error {
	clientID := ""
	username := ""
	password := ""
	// create broker specific connection with wanted params
	brokerConnectionOptions := &mqtt.Connect{
		KeepAlive:  30,
		ClientID:   clientID,
		CleanStart: true,
		Username:   username,
		Password:   []byte(password),
	}
	// create a broker client
	brokerConnection, err := m.mqttInstance.Connect(context.Background(), brokerConnectionOptions)
	if err != nil {
		log.Println("broker connection error instance")
		log.Println(err)
		return err
	}
	if brokerConnection.ReasonCode != 0 {
		log.Println("Failed to connect to MQTT Broker")
		return err
	}
	return nil
}

func (m *MqttClient) Pubblish(topic string, message []byte, qos uint8) error {
	packet := &mqtt.Publish{
		Topic:   topic,
		QoS:     qos,
		Retain:  false,
		Payload: message,
	}
	_, err := m.mqttInstance.Publish(context.Background(), packet)
	if err != nil {
		return err
	}
	return nil
}
