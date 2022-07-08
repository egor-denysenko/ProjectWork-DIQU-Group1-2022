package mqttaccess

import (
	"context"
	mqtt "github.com/eclipse/paho.golang/paho"
	"log"
	"net"
	"os"
)

type BrokerSubscriptions *mqtt.Subscribe

type MqttClient struct {
	mqttInstance        *mqtt.Client
	BrokerSubscriptions BrokerSubscriptions
}

func NewMqttConnection(mqttSubChan chan<- []byte) *MqttClient {
	serverUrl := os.Getenv("MqttBrokerAdd")
	// Try to reach the Broker and connects to it
	brokerDial, err := net.Dial("tcp4", serverUrl)
	// manage the connection error from the Dial
	if err != nil {
		log.Fatalf("Failed to connect to %v", err)
	}
	// Create a default broker instance
	brokerClient := mqtt.NewClient(mqtt.ClientConfig{
		Router: mqtt.NewSingleHandlerRouter(func(m *mqtt.Publish) {
			mqttSubChan <- m.Payload
		}),
		Conn: brokerDial,
	})

	initializeBrokerSubscriptions := mqtt.Subscribe{
		Properties: &mqtt.SubscribeProperties{
			SubscriptionIdentifier: nil,
			User:                   nil,
		},
		Subscriptions: map[string]mqtt.SubscribeOptions{},
	}

	return &MqttClient{
		mqttInstance:        brokerClient,
		BrokerSubscriptions: &initializeBrokerSubscriptions,
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
		log.Println(err)
		log.Printf("%T %+v", err, err)
		return err
	}
	if brokerConnection.ReasonCode != 0 {
		log.Fatalf("Failed to connect to MQTT Broker")
		return err
	}
	return nil
}

func (m *MqttClient) Subscribe(ctx context.Context, topic string) error {
	m.BrokerSubscriptions.Subscriptions[topic] = mqtt.SubscribeOptions{
		QoS:               0,
		RetainHandling:    0,
		NoLocal:           false,
		RetainAsPublished: false,
	}
	//m.BrokerSubscriptions.Properties
	subAck, err := m.mqttInstance.Subscribe(ctx, m.BrokerSubscriptions)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(*subAck)
	return nil
}
