package mqttservice

import (
	"MqttReceiver/pkg/mqttservice/mqttaccess"
	"MqttReceiver/pkg/mqttservice/mqttlogic"
	"context"
)

type MqttService struct {
	service *mqttlogic.MqttLogic
}

func FactoryMqttService(mqttSubChan chan<- []byte) *MqttService {
	return &MqttService{service: mqttlogic.FactoryMqttLogic(mqttaccess.NewMqttConnection(mqttSubChan))}
}

func (m *MqttService) Connect() error {
	return m.service.Connect()
}
func (m *MqttService) Subscribe(ctx context.Context, topic string) error {
	return m.service.Subscribe(ctx, topic)
}
